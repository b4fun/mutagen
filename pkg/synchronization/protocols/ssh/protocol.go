package ssh

import (
	"context"
	"fmt"
	"io"

	"github.com/mutagen-io/mutagen/pkg/agent"
	"github.com/mutagen-io/mutagen/pkg/agent/transport/ssh"
	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/synchronization"
	"github.com/mutagen-io/mutagen/pkg/synchronization/endpoint/remote"
	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
)

// NewSSHTransportFunc defines the SSH transport creation function type.
type NewSSHTransportFunc func(user, host string, port uint16, prompter string) (agent.Transport, error)

// HandlerOpts configures the SSH protocol handler.
type HandlerOpts interface {
	// apply applies the configurations to the protocol handler.
	apply(*protocolHandler)
}

type handlerOptsFunc func(*protocolHandler)

func (f handlerOptsFunc) apply(h *protocolHandler) {
	f(h)
}

// WithSSHTransportFunc configures the SSH protocol handler to use the specified
// transport creation function.
func WithSSHTransportFunc(f NewSSHTransportFunc) HandlerOpts {
	return handlerOptsFunc(func(h *protocolHandler) {
		h.newTransportFunc = f
	})
}

// WithAgentDialFunc configures the SSH protocol handler to use the specified
// agent dial function.
func WithAgentDialFunc(f agent.DialFunc) HandlerOpts {
	return handlerOptsFunc(func(h *protocolHandler) {
		h.dialFunc = f
	})
}

// protocolHandler implements the synchronization.ProtocolHandler interface for
// connecting to remote endpoints over SSH. It uses the agent infrastructure
// over an SSH transport.
type protocolHandler struct {
	newTransportFunc NewSSHTransportFunc
	dialFunc         agent.DialFunc
}

// dialResult provides asynchronous agent dialing results.
type dialResult struct {
	// stream is the stream returned by agent dialing.
	stream io.ReadWriteCloser
	// error is the error returned by agent dialing.
	error error
}

func (h *protocolHandler) newTransport(
	user, host string,
	port uint16,
	prompter string,
) (agent.Transport, error) {
	newFunc := h.newTransportFunc
	if newFunc == nil {
		newFunc = ssh.NewTransport
	}

	return newFunc(user, host, port, prompter)
}

func (h *protocolHandler) dialAgent(
	logger *logging.Logger, transport agent.Transport, mode, prompter string,
) (io.ReadWriteCloser, error) {
	dialFunc := h.dialFunc
	if dialFunc == nil {
		dialFunc = agent.Dial
	}

	return dialFunc(logger, transport, mode, prompter)
}

// Connect connects to an SSH endpoint.
func (h *protocolHandler) Connect(
	ctx context.Context,
	logger *logging.Logger,
	url *urlpkg.URL,
	prompter string,
	session string,
	version synchronization.Version,
	configuration *synchronization.Configuration,
	alpha bool,
) (synchronization.Endpoint, error) {
	// Verify that the URL is of the correct kind and protocol.
	if url.Kind != urlpkg.Kind_Synchronization {
		panic("non-synchronization URL dispatched to synchronization protocol handler")
	} else if url.Protocol != urlpkg.Protocol_SSH {
		panic("non-SSH URL dispatched to SSH protocol handler")
	}

	// Create an SSH agent transport.
	transport, err := h.newTransport(url.User, url.Host, uint16(url.Port), prompter)
	if err != nil {
		return nil, fmt.Errorf("unable to create SSH transport: %w", err)
	}

	// Create a channel to deliver the dialing result.
	results := make(chan dialResult)

	// Perform dialing in a background Goroutine so that we can monitor for
	// cancellation.
	go func() {
		// Perform the dialing operation.
		stream, err := h.dialAgent(logger, transport, agent.CommandSynchronizer, prompter)

		// Transmit the result or, if cancelled, close the stream.
		select {
		case results <- dialResult{stream, err}:
		case <-ctx.Done():
			if stream != nil {
				stream.Close()
			}
		}
	}()

	// Wait for dialing results or cancellation.
	var stream io.ReadWriteCloser
	select {
	case result := <-results:
		if result.error != nil {
			return nil, fmt.Errorf("unable to dial agent endpoint: %w", result.error)
		}
		stream = result.stream
	case <-ctx.Done():
		return nil, context.Canceled
	}

	// Create the endpoint client.
	return remote.NewEndpoint(logger, stream, url.Path, session, version, configuration, alpha)
}

func init() {
	Register()
}

// Register registers the SSH protocol handler with the synchronization package.
func Register(opts ...HandlerOpts) {
	handler := &protocolHandler{}
	for _, opt := range opts {
		opt.apply(handler)
	}

	synchronization.ProtocolHandlers[urlpkg.Protocol_SSH] = handler
}
