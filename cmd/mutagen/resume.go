package main

import (
	"context"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/havoc-io/mutagen/cmd"
	promptpkg "github.com/havoc-io/mutagen/pkg/prompt"
	sessionsvcpkg "github.com/havoc-io/mutagen/pkg/session/service"
)

func resumeMain(command *cobra.Command, arguments []string) {
	// Parse session specification.
	var sessionQueries []string
	if len(arguments) > 0 {
		if resumeConfiguration.all {
			cmd.Fatal(errors.New("-a/--all specified with specific sessions"))
		}
		sessionQueries = arguments
	} else if !resumeConfiguration.all {
		cmd.Fatal(errors.New("no sessions specified"))
	}

	// Connect to the daemon and defer closure of the connection.
	daemonConnection, err := createDaemonClientConnection()
	if err != nil {
		cmd.Fatal(errors.Wrap(err, "unable to connect to daemon"))
	}
	defer daemonConnection.Close()

	// Create a session service client.
	sessionService := sessionsvcpkg.NewSessionClient(daemonConnection)

	// Invoke the session resume method. The stream will close when the
	// associated context is cancelled.
	resumeContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := sessionService.Resume(resumeContext)
	if err != nil {
		cmd.Fatal(errors.Wrap(err, "unable to invoke resume"))
	}

	// Send the initial request.
	request := &sessionsvcpkg.ResumeRequest{
		All:            resumeConfiguration.all,
		SessionQueries: sessionQueries,
	}
	if err := stream.Send(request); err != nil {
		cmd.Fatal(errors.Wrap(err, "unable to send resume request"))
	}

	// Receive and process responses until we're done.
	for {
		// Receive the next response, watching for completion or another prompt.
		var prompt *promptpkg.Prompt
		if response, err := stream.Recv(); err != nil {
			cmd.Fatal(errors.Wrap(err, "unable to receive response"))
		} else if response.Prompt == nil {
			return
		} else {
			prompt = response.Prompt
		}

		// Process the prompt.
		if response, err := promptpkg.PromptCommandLine(prompt.Message, prompt.Prompt); err != nil {
			cmd.Fatal(errors.Wrap(err, "unable to perform prompting"))
		} else if err = stream.Send(&sessionsvcpkg.ResumeRequest{Response: response}); err != nil {
			cmd.Fatal(errors.Wrap(err, "unable to send prompt response"))
		}
	}
}

var resumeCommand = &cobra.Command{
	Use:   "resume [<session>...]",
	Short: "Resumes a paused or disconnected synchronization session",
	Run:   resumeMain,
}

var resumeConfiguration struct {
	all  bool
	help bool
}

func init() {
	// Bind flags to configuration. We manually add help to override the default
	// message, but Cobra still implements it automatically.
	flags := resumeCommand.Flags()
	flags.BoolVarP(&resumeConfiguration.all, "all", "a", false, "Resume all sessions")
	flags.BoolVarP(&resumeConfiguration.help, "help", "h", false, "Show help information")
}
