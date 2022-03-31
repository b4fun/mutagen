package local

import (
	"context"
	"errors"
	"fmt"
	"hash"
	"io"
	"path/filepath"
	"sync"
	"time"

	"github.com/mutagen-io/mutagen/pkg/encoding"
	"github.com/mutagen-io/mutagen/pkg/filesystem"
	"github.com/mutagen-io/mutagen/pkg/filesystem/behavior"
	"github.com/mutagen-io/mutagen/pkg/filesystem/watching"
	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/sidecar"
	"github.com/mutagen-io/mutagen/pkg/synchronization"
	"github.com/mutagen-io/mutagen/pkg/synchronization/core"
	"github.com/mutagen-io/mutagen/pkg/synchronization/rsync"
)

const (
	// cacheSaveInterval is the interval at which caches are serialized and
	// written to disk in the background.
	cacheSaveInterval = 60 * time.Second
	// recheckPathsMaximumCapacity is the maximum re-check path set capacity.
	recheckPathsMaximumCapacity = 10 * 1024
)

// reifiedWatchMode describes a fully reified watch mode based on the watch mode
// specified for the endpoint and the availability of modes on the system.
type reifiedWatchMode uint8

const (
	// reifiedWatchModeDisabled indicates that watching has been disabled.
	reifiedWatchModeDisabled reifiedWatchMode = iota
	// reifiedWatchModePoll indicates poll-based watching is in use.
	reifiedWatchModePoll
	// reifiedWatchModeRecursive indicates that recursive watching is in use.
	reifiedWatchModeRecursive
)

// endpoint provides a local, in-memory implementation of
// synchronization.Endpoint for local files.
type endpoint struct {
	// logger is the underlying logger. This field is static and thus safe for
	// concurrent usage.
	logger *logging.Logger
	// root is the synchronization root. This field is static and thus safe for
	// concurrent reads.
	root string
	// readOnly determines whether or not the endpoint should be operating in a
	// read-only mode (i.e. it is the source of unidirectional synchronization).
	// This field is static and thus safe for concurrent reads.
	readOnly bool
	// maximumEntryCount is the maximum number of entries that the endpoint will
	// synchronize. This field is static and thus safe for concurrent reads.
	maximumEntryCount uint64
	// watchMode indicates the watch mode being used. This field is static and
	// thus safe for concurrent reads.
	watchMode reifiedWatchMode
	// accelerationAllowed indicates whether or not scan acceleration is
	// allowed. This field is static and thus safe for concurrent reads.
	accelerationAllowed bool
	// probeMode is the probe mode. This field is static and thus safe for
	// concurrent reads.
	probeMode behavior.ProbeMode
	// symbolicLinkMode is the symbolic link mode. This field is static and thus
	// safe for concurrent reads.
	symbolicLinkMode core.SymbolicLinkMode
	// ignores are the path ignore specifications. This field is static and thus
	// safe for concurrent reads.
	ignores []string
	// defaultFileMode is the default file permission mode to use in "portable"
	// permission propagation. This field is static and thus safe for concurrent
	// reads.
	defaultFileMode filesystem.Mode
	// defaultDirectoryMode is the default directory permission mode to use in
	// "portable" permission propagation. This field is static and thus safe for
	// concurrent reads.
	defaultDirectoryMode filesystem.Mode
	// defaultOwnership is the default ownership specification to use in
	// "portable" permission propagation. This field is static and thus safe for
	// concurrent reads.
	defaultOwnership *filesystem.OwnershipSpecification
	// workerCancel cancels any background worker Goroutines for the endpoint.
	// This field is static and thus safe for concurrent invocation.
	workerCancel context.CancelFunc
	// saveCacheDone is closed when the cache saving Goroutine has completed. It
	// will never have values written to it and will only be closed, so a
	// receive that returns indicates closure. This field is static and thus
	// safe for concurrent receive operations.
	saveCacheDone <-chan struct{}
	// watchDone is closed when the watching Goroutine has completed. It will
	// never have values written to it and will only be closed, so a receive
	// that returns indicates closure. This field is static and thus safe for
	// concurrent receive operations.
	watchDone <-chan struct{}
	// pollEvents is the channel used to inform a call to Poll that there are
	// filesystem modifications (and thus it can return). It is a buffered
	// channel with a capacity of one. Senders should always perform a
	// non-blocking send to the channel, because if it is already populated,
	// then filesystem modifications are already indicated. This field is static
	// and never closed, and is thus safe for concurrent send operations.
	pollEvents chan struct{}
	// recursiveWatchRetryEstablish is a channel used by Transition to signal to
	// the recursive watching Goroutine (if any) that it should try to
	// re-establish watching. It is a non-buffered channel, with reads only
	// occurring when the recursive watching Goroutine is waiting to retry watch
	// establishment and writes only occurring in a non-blocking fashion
	// (meaning this is a best-effort signaling mechanism (with a fallback to a
	// timer-based signal)). This field is static and never closed, and is thus
	// safe for concurrent send operations.
	recursiveWatchRetryEstablish chan struct{}
	// scanLock serializes access to accelerate, recheckPaths, snapshot, hasher,
	// cache, ignoreCache, cacheWriteError, lastScanEntryCount,
	// scannedSinceLastStageCall, and scannedSinceLastTransitionCall. This lock
	// is not necessitated by the Endpoint interface (which doesn't permit
	// concurrent usage), but rather the endpoint's background worker Goroutines
	// for cache saving and filesystem watching.
	scanLock sync.Mutex
	// accelerate indicates that the Scan function should attempt to accelerate
	// scanning by using data from a background watcher Goroutine.
	accelerate bool
	// recheckPaths is the set of re-check paths to use when accelerating scans
	// in recursive watching mode. This map will be non-nil if and only if
	// accelerate is true and recursive watching is being used.
	recheckPaths map[string]bool
	// snapshot is the snapshot from the last scan.
	snapshot *core.Snapshot
	// hasher is the hasher used for scans.
	hasher hash.Hash
	// cache is the cache from the last successful scan on the endpoint.
	cache *core.Cache
	// ignoreCache is the ignore cache from the last successful scan on the
	// endpoint.
	ignoreCache core.IgnoreCache
	// cacheWriteError is the last error encountered when trying to write the
	// cache to disk, if any.
	cacheWriteError error
	// lastScanEntryCount is the entry count at the time of the last scan.
	lastScanEntryCount uint64
	// scannedSinceLastStageCall tracks whether or not a scan operation has
	// occurred since the last staging operation.
	scannedSinceLastStageCall bool
	// scannedSinceLastTransitionCall tracks whether or not a scan operation has
	// occurred since the last transitioning operation.
	scannedSinceLastTransitionCall bool
	// stager is the staging coordinator. It is not safe for concurrent usage,
	// but since Endpoint doesn't allow concurrent usage, we know that the
	// stager will only be used in at most one of Stage or Transition methods at
	// any given time.
	stager *stager
}

// NewEndpoint creates a new local endpoint instance using the specified session
// metadata and options.
func NewEndpoint(
	logger *logging.Logger,
	root string,
	sessionIdentifier string,
	version synchronization.Version,
	configuration *synchronization.Configuration,
	alpha bool,
) (synchronization.Endpoint, error) {
	// Determine if the endpoint is running in a read-only mode.
	synchronizationMode := configuration.SynchronizationMode
	if synchronizationMode.IsDefault() {
		synchronizationMode = version.DefaultSynchronizationMode()
	}
	unidirectional := synchronizationMode == core.SynchronizationMode_SynchronizationModeOneWaySafe ||
		synchronizationMode == core.SynchronizationMode_SynchronizationModeOneWayReplica
	readOnly := alpha && unidirectional

	// Determine the maximum entry count.
	maximumEntryCount := configuration.MaximumEntryCount
	if maximumEntryCount == 0 {
		maximumEntryCount = version.DefaultMaximumEntryCount()
	}

	// Determine the maximum staging file size.
	maximumStagingFileSize := configuration.MaximumStagingFileSize
	if maximumStagingFileSize == 0 {
		maximumStagingFileSize = version.DefaultMaximumStagingFileSize()
	}

	// Compute the effective watch mode.
	watchMode := configuration.WatchMode
	if watchMode.IsDefault() {
		watchMode = version.DefaultWatchMode()
	}

	// Compute the actual (reified) watch mode.
	var actualWatchMode reifiedWatchMode
	var nonRecursiveWatchingAllowed bool
	if watchMode == synchronization.WatchMode_WatchModePortable {
		if watching.RecursiveWatchingSupported {
			actualWatchMode = reifiedWatchModeRecursive
		} else {
			actualWatchMode = reifiedWatchModePoll
			nonRecursiveWatchingAllowed = true
		}
	} else if watchMode == synchronization.WatchMode_WatchModeForcePoll {
		actualWatchMode = reifiedWatchModePoll
	} else if watchMode == synchronization.WatchMode_WatchModeNoWatch {
		actualWatchMode = reifiedWatchModeDisabled
	} else {
		panic("unhandled watch mode")
	}

	// Compute the effective scan mode and determine whether or not scan
	// acceleration is allowed.
	scanMode := configuration.ScanMode
	if scanMode.IsDefault() {
		scanMode = version.DefaultScanMode()
	}
	accelerationAllowed := scanMode == synchronization.ScanMode_ScanModeAccelerated

	// Compute the effective probe mode.
	probeMode := configuration.ProbeMode
	if probeMode.IsDefault() {
		probeMode = version.DefaultProbeMode()
	}

	// Compute the effective symbolic link mode.
	symbolicLinkMode := configuration.SymbolicLinkMode
	if symbolicLinkMode.IsDefault() {
		symbolicLinkMode = version.DefaultSymbolicLinkMode()
	}

	// Compute the effective VCS ignore mode.
	ignoreVCSMode := configuration.IgnoreVCSMode
	if ignoreVCSMode.IsDefault() {
		ignoreVCSMode = version.DefaultIgnoreVCSMode()
	}

	// Compute a combined ignore list.
	var ignores []string
	if ignoreVCSMode == core.IgnoreVCSMode_IgnoreVCSModeIgnore {
		ignores = append(ignores, core.DefaultVCSIgnores...)
	}
	ignores = append(ignores, configuration.DefaultIgnores...)
	ignores = append(ignores, configuration.Ignores...)

	// Track whether or not any non-default ownership or directory permissions
	// are set. We don't care about non-default file permissions since we're
	// only tracking this to set volume root ownership and permissions in
	// sidecar containers.
	var nonDefaultOwnershipOrDirectoryPermissionsSet bool

	// Compute the effective default file mode.
	defaultFileMode := filesystem.Mode(configuration.DefaultFileMode)
	if defaultFileMode == 0 {
		defaultFileMode = version.DefaultFileMode()
	}

	// Compute the effective default directory mode.
	defaultDirectoryMode := filesystem.Mode(configuration.DefaultDirectoryMode)
	if defaultDirectoryMode == 0 {
		defaultDirectoryMode = version.DefaultDirectoryMode()
	} else {
		nonDefaultOwnershipOrDirectoryPermissionsSet = true
	}

	// Compute the effective owner specification.
	defaultOwnerSpecification := configuration.DefaultOwner
	if defaultOwnerSpecification == "" {
		defaultOwnerSpecification = version.DefaultOwnerSpecification()
	} else {
		nonDefaultOwnershipOrDirectoryPermissionsSet = true
	}

	// Compute the effective owner group specification.
	defaultGroupSpecification := configuration.DefaultGroup
	if defaultGroupSpecification == "" {
		defaultGroupSpecification = version.DefaultGroupSpecification()
	} else {
		nonDefaultOwnershipOrDirectoryPermissionsSet = true
	}

	// Compute the effective ownership specification.
	defaultOwnership, err := filesystem.NewOwnershipSpecification(
		defaultOwnerSpecification,
		defaultGroupSpecification,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create ownership specification: %w", err)
	}

	// Compute the cache path if this isn't an ephemeral endpoint.
	cachePath, err := pathForCache(sessionIdentifier, alpha)
	if err != nil {
		return nil, fmt.Errorf("unable to compute/create cache path: %w", err)
	}

	// Load any existing cache. If it fails to load or validate, just replace it
	// with an empty one.
	// TODO: Should we let validation errors bubble up? They may be indicative
	// of something bad.
	cache := &core.Cache{}
	if encoding.LoadAndUnmarshalProtobuf(cachePath, cache) != nil {
		cache = &core.Cache{}
	} else if cache.EnsureValid() != nil {
		cache = &core.Cache{}
	}

	// Check if the synchronization root is a volume Mount point in a Mutagen
	// sidecar container.
	var rootIsSidecarVolumeMountPoint bool
	var sidecarVolumeName string
	if sidecar.EnvironmentIsSidecar() {
		rootIsSidecarVolumeMountPoint, sidecarVolumeName = sidecar.PathIsVolumeMountPoint(root)
	}

	// Compute the effective staging mode. If no mode has been explicitly set
	// and the synchronization root is a volume mount point in a Mutagen sidecar
	// container, then use internal staging for better performance. Otherwise,
	// use either the explicitly specified staging mode or the default staging
	// mode.
	stageMode := configuration.StageMode
	if stageMode.IsDefault() {
		if rootIsSidecarVolumeMountPoint {
			stageMode = synchronization.StageMode_StageModeInternal
		} else {
			stageMode = version.DefaultStageMode()
		}
	}

	// Compute the staging root path and whether or not it should be hidden.
	var stagingRoot string
	var hideStagingRoot bool
	if stageMode == synchronization.StageMode_StageModeMutagen {
		stagingRoot, err = pathForMutagenStagingRoot(sessionIdentifier, alpha)
	} else if stageMode == synchronization.StageMode_StageModeNeighboring {
		stagingRoot, err = pathForNeighboringStagingRoot(root, sessionIdentifier, alpha)
		hideStagingRoot = true
	} else if stageMode == synchronization.StageMode_StageModeInternal {
		stagingRoot, err = pathForInternalStagingRoot(root, sessionIdentifier, alpha)
		hideStagingRoot = true
	} else {
		panic("unhandled staging mode")
	}
	if err != nil {
		return nil, fmt.Errorf("unable to compute staging root: %w", err)
	}

	// HACK: If non-default ownership or permissions have been set and the
	// synchronization root is a volume mount point in a Mutagen sidecar
	// container with no pre-existing content, then set the ownership and
	// permissions of the synchronization root to match those of the session.
	// This is a heuristic to work around the fact that Docker volumes don't
	// allow ownership specification at creation time, either via the command
	// line or Compose.
	if nonDefaultOwnershipOrDirectoryPermissionsSet && rootIsSidecarVolumeMountPoint {
		if err := sidecar.SetVolumeOwnershipAndPermissionsIfEmpty(
			sidecarVolumeName,
			defaultOwnership,
			defaultDirectoryMode,
		); err != nil {
			return nil, fmt.Errorf("unable to set ownership and permissions for sidecar volume: %w", err)
		}
	}

	// Create a cancellable context in which the endpoint's background worker
	// Goroutines will operate.
	workerContext, workerCancel := context.WithCancel(context.Background())

	// Create channels to monitor background worker Goroutine completion.
	saveCacheDone := make(chan struct{})
	watchDone := make(chan struct{})

	// Create the endpoint.
	endpoint := &endpoint{
		logger:                       logger,
		root:                         root,
		readOnly:                     readOnly,
		maximumEntryCount:            maximumEntryCount,
		watchMode:                    actualWatchMode,
		accelerationAllowed:          accelerationAllowed,
		probeMode:                    probeMode,
		symbolicLinkMode:             symbolicLinkMode,
		ignores:                      ignores,
		defaultFileMode:              defaultFileMode,
		defaultDirectoryMode:         defaultDirectoryMode,
		defaultOwnership:             defaultOwnership,
		workerCancel:                 workerCancel,
		saveCacheDone:                saveCacheDone,
		watchDone:                    watchDone,
		pollEvents:                   make(chan struct{}, 1),
		recursiveWatchRetryEstablish: make(chan struct{}),
		hasher:                       version.Hasher(),
		cache:                        cache,
		stager: newStager(
			stagingRoot,
			hideStagingRoot,
			version.Hasher(),
			maximumStagingFileSize,
		),
	}

	// Start the cache saving Goroutine.
	go func() {
		endpoint.saveCacheRegularly(workerContext, cachePath)
		close(saveCacheDone)
	}()

	// Compute the effective watch polling interval.
	watchPollingInterval := configuration.WatchPollingInterval
	if watchPollingInterval == 0 {
		watchPollingInterval = version.DefaultWatchPollingInterval()
	}

	// Start the watching Goroutine.
	go func() {
		if actualWatchMode == reifiedWatchModePoll {
			endpoint.watchPoll(workerContext, watchPollingInterval, nonRecursiveWatchingAllowed)
		} else if actualWatchMode == reifiedWatchModeRecursive {
			go endpoint.watchRecursive(workerContext, watchPollingInterval)
		}
		close(watchDone)
	}()

	// Success.
	return endpoint, nil
}

// saveCacheRegularly serializes the cache and writes the result to disk at
// regular intervals. It runs as a background Goroutine for all endpoints.
func (e *endpoint) saveCacheRegularly(context context.Context, cachePath string) {
	// Create a ticker to regulate cache saving and defer its shutdown.
	ticker := time.NewTicker(cacheSaveInterval)
	defer ticker.Stop()

	// Track the last saved cache. If it hasn't changed, there's no point in
	// rewriting it. It's safe to keep a reference to the cache since caches are
	// treated as immutable. The only cost is keeping an old cache around until
	// the next write cycle, but that's a relatively small price to pay to avoid
	// unnecessary disk writes.
	var lastSavedCache *core.Cache

	// Loop indefinitely, watching for cancellation and saving the cache to
	// disk at regular intervals. If we see a cache write failure, we record it,
	// and we don't attempt any more saves. The recorded error will be reported
	// to the controller on the next call to Scan.
	for {
		select {
		case <-context.Done():
			return
		case <-ticker.C:
			e.scanLock.Lock()
			if e.cacheWriteError == nil && e.cache != lastSavedCache {
				if err := encoding.MarshalAndSaveProtobuf(cachePath, e.cache); err != nil {
					e.cacheWriteError = err
				} else {
					lastSavedCache = e.cache
				}
			}
			e.scanLock.Unlock()
		}
	}
}

// stopAndDrainTimer stops a timer and performs a non-blocking drain on its
// channel. This allows a timer to be stopped and drained without any knowledge
// of its current state.
func stopAndDrainTimer(timer *time.Timer) {
	timer.Stop()
	select {
	case <-timer.C:
	default:
	}
}

// watchPoll is the watch loop for poll-based watching, with optional support
// for using native non-recursive watching facilities to reduce notification
// latency on frequently updated contents.
func (e *endpoint) watchPoll(ctx context.Context, pollingInterval uint32, nonRecursiveWatchingAllowed bool) {
	// Create a sublogger.
	logger := e.logger.Sublogger("polling")

	// Create a ticker to regulate polling and defer its shutdown.
	ticker := time.NewTicker(time.Duration(pollingInterval) * time.Second)
	defer ticker.Stop()

	// Track whether or not it's our first iteration in the polling loop. We
	// adjust some behaviors in that case.
	first := true

	// Track the previous snapshot.
	previous := &core.Snapshot{}

	// If non-recursive watching is available, then set up a non-recursive
	// watcher. Since non-recursive watching is a best-effort basis to reduce
	// latency, we don't try to re-establish this watcher if it fails.
	var watcher watching.NonRecursiveWatcher
	var watchEvents <-chan map[string]bool
	var watchErrors <-chan error
	if nonRecursiveWatchingAllowed && watching.NonRecursiveWatchingSupported {
		// Create the filter that we'll use to exclude Mutagen temporary files.
		filter := func(path string) bool {
			return filesystem.IsTemporaryFileName(filepath.Base(path))
		}

		// Attempt to create the watcher and ensure that it will be terminated.
		logger.Debug("Creating non-recursive watcher")
		if w, err := watching.NewNonRecursiveWatcher(filter); err != nil {
			logger.Debug("Unable to create non-recursive watcher:", err)
		} else {
			logger.Debug("Successfully created non-recursive watcher")
			watcher = w
			watchEvents = watcher.Events()
			watchErrors = watcher.Errors()
			defer func() {
				if watcher != nil {
					watcher.Terminate()
				}
			}()
		}
	}

	// Loop until cancellation, performing polling at the specified interval.
	for {
		// Set behaviors based on whether or not this is our first time in the
		// loop. If this is our first time in the loop, then we skip waiting,
		// because our ticker won't fire its first event until after the polling
		// duration has elapsed, and we'd like a baseline scan before that. The
		// reason we want a baseline scan before that is that we'll ignore
		// modifications on our first successful scan. The reason for ignoring
		// these modifications is that we'll be comparing against zero-valued
		// variables and are thus certain to see modifications if there is
		// existing content on disk. Since the controller already skips polling
		// (if watching is enabled) on its first synchronization cycle, there's
		// no point for us to also send a notification, because if both
		// endpoints did this, you'd see up to three scans on session startup.
		// Of course, if our scan fails on the first try, then we'll allow a
		// notification (due to these "artificial" modifications) to be sent
		// after the first successful scan, but that will at least occur after
		// the initial polling duration.
		var skipWaiting, ignoreModifications bool
		if first {
			skipWaiting = true
			ignoreModifications = true
			first = false
		}

		// Unless we're skipping waiting, wait for cancellation, a tick event,
		// a notification from our non-recursive watches, or a coalesced event.
		if !skipWaiting {
			select {
			case <-ctx.Done():
				// Log termination.
				logger.Debug("Polling terminated")

				// Ensure that accelerated watching is disabled, if necessary.
				if e.accelerationAllowed {
					e.scanLock.Lock()
					e.accelerate = false
					e.scanLock.Unlock()
				}

				// Terminate polling.
				return
			case <-ticker.C:
				logger.Trace("Received polling signal")
			case err := <-watchErrors:
				// Log the error.
				logger.Debug("Non-recursive watching error:", err)

				// Terminate the watcher and nil it out. We don't bother trying
				// to re-establish it. Also nil out the errors channel in case
				// the watcher pumps any additional errors into it (in which
				// case we don't want to trigger this code again on a nil
				// watcher). We'll allow event channels to continue since they
				// may contain residual events.
				watcher.Terminate()
				watcher = nil
				watchErrors = nil

				// Continue polling.
				continue
			case event := <-watchEvents:
				logger.Trace("Received event with", len(event), "paths")
			}
		}

		// Grab the scan lock.
		e.scanLock.Lock()

		// Disable the use of the existing scan results.
		e.accelerate = false

		// Perform a scan. If there's an error, then assume it's due to
		// concurrent modification. In that case, release the scan lock and
		// strobe the poll events channel. The controller can then perform a
		// full scan.
		if err := e.scan(ctx, nil, nil); err != nil {
			// Log the error.
			logger.Debug("Scan failed:", err)

			// Release the scan lock.
			e.scanLock.Unlock()

			// Strobe the poll events channel and continue polling.
			e.strobePollEvents()
			continue
		}

		// If our scan was successful, then we know that the scan results
		// will be okay to return for the next Scan call, though we only
		// indicate that acceleration should be used if the endpoint allows it.
		e.accelerate = e.accelerationAllowed

		// Extract scan parameters so that we can release the scan lock.
		snapshot := e.snapshot

		// Release the scan lock.
		e.scanLock.Unlock()

		// Check for modifications.
		modified := !snapshot.Equal(previous)

		// If we have a working non-recursive watcher, then perform a full diff
		// to determine new watch paths, and then start the new watches. Any
		// watch errors will be reported on the watch errors channel.
		if watcher != nil {
			changes := core.Diff(previous.Content, snapshot.Content)
			for _, change := range changes {
				watcher.Watch(filepath.Join(e.root, change.Path))
			}
		}

		// Update our tracking parameters.
		previous = snapshot

		// If we've seen modifications, and we're not ignoring them, then strobe
		// the poll events channel.
		if modified && !ignoreModifications {
			// Log the modifications.
			logger.Trace("Modifications detected")

			// Strobe the poll events channel.
			e.strobePollEvents()
		} else {
			// Log the lack of modifications.
			logger.Trace("No modifications detected")
		}
	}
}

// watchRecursive is the watch loop for platforms where native recursive
// watching facilities are available.
func (e *endpoint) watchRecursive(ctx context.Context, pollingInterval uint32) {
	// Create a sublogger.
	logger := e.logger.Sublogger("watching")

	// Convert the polling interval to a duration.
	pollingDuration := time.Duration(pollingInterval) * time.Second

	// Track our recursive watcher and ensure that it's stopped when we return.
	var watcher watching.RecursiveWatcher
	defer func() {
		if watcher != nil {
			watcher.Terminate()
		}
	}()

	// Create the filter that we'll use to exclude Mutagen temporary files.
	// Recursive watchers can use our fast-path base name calculation.
	filter := func(path string) bool {
		return filesystem.IsTemporaryFileName(core.PathBase(path))
	}

	// Create a timer, initially stopped and drained, that we can use to
	// regulate waiting periods. Also, ensure that it's stopped when we return.
	timer := time.NewTimer(0)
	stopAndDrainTimer(timer)
	defer timer.Stop()

	// Loop until cancellation.
	var err error
WatchEstablishment:
	for {
		// Attempt to establish the watch.
		logger.Debug("Attempting to establish recursive watch")
		watcher, err = watching.NewRecursiveWatcher(e.root, filter)
		if err != nil {
			// Log the failure.
			logger.Debug("Unable to establish recursive watch:", err)

			// Strobe poll events (since nothing else will be driving
			// synchronization from this endpoint at this point in time).
			e.strobePollEvents()

			// Wait to retry watch establishment.
			timer.Reset(pollingDuration)
			select {
			case <-ctx.Done():
				logger.Debug("Watching terminated while waiting for establishment")
				return
			case <-timer.C:
				continue
			case <-e.recursiveWatchRetryEstablish:
				logger.Debug("Received recursive watch establishment suggestion")
				stopAndDrainTimer(timer)
				continue
			}
		}
		logger.Debug("Watch successfully established")

		// Strobe the poll events channel to signal that we now know the
		// synchronization root exists and is accessible.
		e.strobePollEvents()

		// If accelerated scanning is allowed, then reset the timer (which won't
		// be running) to fire immediately in the event loop in order to try
		// enabling acceleration.
		if e.accelerationAllowed {
			timer.Reset(0)
		}

		// Loop and process events.
		for {
			select {
			case <-ctx.Done():
				// Log termination.
				logger.Debug("Watching terminated")

				// Ensure that accelerated watching is disabled, if necessary.
				if e.accelerationAllowed {
					e.scanLock.Lock()
					e.accelerate = false
					e.recheckPaths = nil
					e.scanLock.Unlock()
				}

				// Terminate watching.
				return
			case <-timer.C:
				// Log the acceleration attempt.
				logger.Debug("Attempting to enable accelerated scanning")

				// Attempt to perform a baseline scan to enable acceleration.
				e.scanLock.Lock()
				if err := e.scan(ctx, nil, nil); err != nil {
					logger.Debug("Unable to perform baseline scan:", err)
					timer.Reset(pollingDuration)
				} else {
					logger.Debug("Accelerated scanning now available")
					e.accelerate = true
					e.recheckPaths = make(map[string]bool)
				}
				e.scanLock.Unlock()
			case err := <-watcher.Errors():
				// Log the error.
				logger.Debug("Recursive watching error:", err)

				// If acceleration is allowed on the endpoint, then disable scan
				// acceleration and clear out the re-check paths.
				if e.accelerationAllowed {
					e.scanLock.Lock()
					e.accelerate = false
					e.recheckPaths = nil
					e.scanLock.Unlock()
				}

				// Stop and drain the timer, which may be running.
				stopAndDrainTimer(timer)

				// Strobe the poll events channel since something has occurred
				// (likely on disk) that's killed our watch.
				e.strobePollEvents()

				// Terminate the watcher and restart watch establishment.
				watcher.Terminate()
				watcher = nil
				continue WatchEstablishment
			case event := <-watcher.Events():
				// Log the event.
				logger.Trace("Received event with", len(event), "paths")

				// If acceleration is allowed (and available) on the endpoint,
				// then register the event's paths as re-check paths. We only
				// need to do this if acceleration is already available,
				// otherwise we're still in a pre-baseline scan state and don't
				// need to record these events. If we overflow the maximum
				// re-check path set size, then we'll disable acceleration and
				// schedule an immediate scan to re-enable acceleration.
				if e.accelerationAllowed {
					e.scanLock.Lock()
					if e.accelerate {
						for path := range event {
							e.recheckPaths[path] = true
							if len(e.recheckPaths) > recheckPathsMaximumCapacity {
								logger.Debug("Re-check paths overflowed maximum capacity")
								e.accelerate = false
								e.recheckPaths = nil
								timer.Reset(0)
								break
							}
						}
					}
					e.scanLock.Unlock()
				}

				// Strobe the poll events channel to signal the event.
				e.strobePollEvents()
			}
		}
	}
}

// strobePollEvents strobes the pollEvents channel in a non-blocking fashion.
func (e *endpoint) strobePollEvents() {
	select {
	case e.pollEvents <- struct{}{}:
	default:
	}
}

// Poll implements the Poll method for local endpoints.
func (e *endpoint) Poll(ctx context.Context) error {
	// Wait for either cancellation or an event.
	select {
	case <-ctx.Done():
	case <-e.pollEvents:
	}

	// Done.
	return nil
}

// scan is the internal function which performs a scan operation on the root and
// updates the endpoint scan parameters. The caller must hold the scan lock.
func (e *endpoint) scan(ctx context.Context, baseline *core.Snapshot, recheckPaths map[string]bool) error {
	// Perform a full (warm) scan, watching for errors.
	snapshot, newCache, newIgnoreCache, err := core.Scan(
		ctx,
		e.root,
		baseline, recheckPaths,
		e.hasher, e.cache,
		e.ignores, e.ignoreCache,
		e.probeMode,
		e.symbolicLinkMode,
	)
	if err != nil {
		return err
	}

	// Update the snapshot.
	e.snapshot = snapshot

	// Update caches.
	e.cache = newCache
	e.ignoreCache = newIgnoreCache

	// Update the last scan entry count.
	e.lastScanEntryCount = snapshot.Content.Count()

	// Update call states.
	e.scannedSinceLastStageCall = true
	e.scannedSinceLastTransitionCall = true

	// Success.
	return nil
}

// Scan implements the Scan method for local endpoints.
func (e *endpoint) Scan(ctx context.Context, _ *core.Entry, full bool) (*core.Snapshot, error, bool) {
	// Grab the scan lock and defer its release.
	e.scanLock.Lock()
	defer e.scanLock.Unlock()

	// Before attempting to perform a scan, check for any cache write errors
	// that may have occurred during background cache writes. If we see any
	// error, then we skip scanning and report them here.
	if e.cacheWriteError != nil {
		return nil, fmt.Errorf("unable to save cache to disk: %w", e.cacheWriteError), false
	}

	// Perform a scan.
	//
	// We check to see if we can accelerate the scanning process by using
	// information from a background watching Goroutine. For recursive watching,
	// this means performing a re-scan using a baseline and a set of re-check
	// paths. For poll-based watching, this just means re-using the last scan,
	// so no action is needed here. If acceleration isn't available (due to the
	// state of the watcher or because it's disallowed on the endpoint), then we
	// just perform a full (warm) scan. We also avoid acceleration in the event
	// that a full scan has been explicitly requested, but we don't make any
	// change to the state of acceleration availability, because performing a
	// full warm scan will only improve the accuracy of the baseline (most
	// recent) snapshot, so acceleration will still work.
	//
	// If we see any error while scanning, we just have to assume that it's due
	// to concurrent modifications and suggest a retry. In the case of
	// accelerated scanning with recursive watching, there's no need to disable
	// acceleration on failure so long as the watch is still established (and if
	// it's not, that will handled elsewhere).
	if e.accelerate && !full {
		if e.watchMode == reifiedWatchModeRecursive {
			if err := e.scan(ctx, e.snapshot, e.recheckPaths); err != nil {
				return nil, err, true
			} else {
				e.recheckPaths = make(map[string]bool)
			}
		}
	} else {
		if err := e.scan(ctx, nil, nil); err != nil {
			return nil, err, true
		}
	}

	// Verify that we haven't exceeded the maximum entry count.
	// TODO: Do we actually want to enforce this count in the scan operation so
	// that we don't hold those entries in memory? Right now this is mostly
	// concerned with avoiding transmission of the entries over the wire.
	if e.lastScanEntryCount > e.maximumEntryCount {
		return nil, errors.New("exceeded allowed entry count"), true
	}

	// Success.
	return e.snapshot, nil, false
}

// stageFromRoot attempts to perform staging from local files by using a reverse
// lookup map.
func (e *endpoint) stageFromRoot(
	path string,
	digest []byte,
	reverseLookupMap *core.ReverseLookupMap,
	opener *filesystem.Opener,
) bool {
	// See if we can find a path within the root that has a matching digest.
	sourcePath, sourcePathOk := reverseLookupMap.Lookup(digest)
	if !sourcePathOk {
		return false
	}

	// Open the source file and defer its closure.
	source, err := opener.OpenFile(sourcePath)
	if err != nil {
		return false
	}
	defer source.Close()

	// Create a staging sink. We explicitly manage its closure below.
	sink, err := e.stager.Sink(path)
	if err != nil {
		return false
	}

	// Copy data to the sink and close it, then check for copy errors.
	_, err = io.Copy(sink, source)
	sink.Close()
	if err != nil {
		return false
	}

	// Ensure that everything staged correctly.
	_, err = e.stager.Provide(path, digest)
	return err == nil
}

// Stage implements the Stage method for local endpoints.
func (e *endpoint) Stage(paths []string, digests [][]byte) ([]string, []*rsync.Signature, rsync.Receiver, error) {
	// If we're in a read-only mode, we shouldn't be staging files.
	if e.readOnly {
		return nil, nil, nil, errors.New("endpoint is in read-only mode")
	}

	// Validate argument lengths and bail if there's nothing to stage.
	if len(paths) != len(digests) {
		return nil, nil, nil, errors.New("path count does not match digest count")
	} else if len(paths) == 0 {
		return nil, nil, nil, nil
	}

	// Grab the scan lock. We'll need this to verify the last scan entry count
	// and to generate the reverse lookup map.
	e.scanLock.Lock()

	// Verify that we've performed a scan since the last staging operation, that
	// way our count check is valid. If we haven't, then the controller is
	// either malfunctioning or malicious.
	if !e.scannedSinceLastStageCall {
		e.scanLock.Unlock()
		return nil, nil, nil, errors.New("multiple staging operations performed without scan")
	}
	e.scannedSinceLastStageCall = false

	// Verify that the number of paths provided isn't going to put us over the
	// maximum number of allowed entries.
	if e.maximumEntryCount != 0 && (e.maximumEntryCount-e.lastScanEntryCount) < uint64(len(paths)) {
		e.scanLock.Unlock()
		return nil, nil, nil, errors.New("staging would exceeded allowed entry count")
	}

	// Generate a reverse lookup map from the cache, which we'll use shortly to
	// detect renames and copies.
	reverseLookupMap, err := e.cache.GenerateReverseLookupMap()
	if err != nil {
		e.scanLock.Unlock()
		return nil, nil, nil, fmt.Errorf("unable to generate reverse lookup map: %w", err)
	}

	// Release the scan lock.
	e.scanLock.Unlock()

	// Create an opener that we can use file opening and defer its closure. We
	// can't cache this across synchronization cycles since its path references
	// may become invalidated or may prevent modifications.
	opener := filesystem.NewOpener(e.root)
	defer opener.Close()

	// Filter the path list by looking for files that we can source locally.
	//
	// First, check if the content can be provided from the stager, which
	// indicates that a previous staging operation was interrupted.
	//
	// Second, use a reverse lookup map (generated from the cache) and see if we
	// can find (and stage) any files locally, which indicates that a file has
	// been copied or renamed.
	//
	// If we manage to handle all files, then we can abort staging.
	filteredPaths := paths[:0]
	for p, path := range paths {
		digest := digests[p]
		if _, err := e.stager.Provide(path, digest); err == nil {
			continue
		} else if e.stageFromRoot(path, digest, reverseLookupMap, opener) {
			continue
		} else {
			filteredPaths = append(filteredPaths, path)
		}
	}
	if len(filteredPaths) == 0 {
		return nil, nil, nil, nil
	}

	// Create an rsync engine.
	engine := rsync.NewEngine()

	// Compute signatures for each of the unstaged paths. For paths that don't
	// exist or that can't be read, just use an empty signature, which means to
	// expect/use an empty base when deltafying/patching.
	signatures := make([]*rsync.Signature, len(filteredPaths))
	for p, path := range filteredPaths {
		if base, err := opener.OpenFile(path); err != nil {
			signatures[p] = &rsync.Signature{}
			continue
		} else if signature, err := engine.Signature(base, 0); err != nil {
			base.Close()
			signatures[p] = &rsync.Signature{}
			continue
		} else {
			base.Close()
			signatures[p] = signature
		}
	}

	// Create a receiver.
	receiver, err := rsync.NewReceiver(e.root, filteredPaths, signatures, e.stager)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("unable to create rsync receiver: %w", err)
	}

	// Done.
	return filteredPaths, signatures, receiver, nil
}

// Supply implements the supply method for local endpoints.
func (e *endpoint) Supply(paths []string, signatures []*rsync.Signature, receiver rsync.Receiver) error {
	return rsync.Transmit(e.root, paths, signatures, receiver)
}

// Transition implements the Transition method for local endpoints.
func (e *endpoint) Transition(ctx context.Context, transitions []*core.Change) ([]*core.Entry, []*core.Problem, bool, error) {
	// If we're in a read-only mode, we shouldn't be performing transitions.
	if e.readOnly {
		return nil, nil, false, errors.New("endpoint is in read-only mode")
	}

	// Grab the scan lock and defer its release.
	e.scanLock.Lock()
	defer e.scanLock.Unlock()

	// Verify that we've performed a scan since the last transition operation,
	// that way our count check is valid. If we haven't, then the controller is
	// either malfunctioning or malicious.
	if !e.scannedSinceLastTransitionCall {
		return nil, nil, false, errors.New("multiple transition operations performed without scan")
	}
	e.scannedSinceLastTransitionCall = false

	// Verify that the number of entries we'll be creating won't put us over the
	// maximum number of allowed entries. Again, we don't worry too much about
	// overflow here for the same reasons as in Entry.Count.
	if e.maximumEntryCount != 0 {
		// Compute the resulting entry count. If we dip below zero in this
		// counting process, then the controller is malfunctioning.
		resultingEntryCount := e.lastScanEntryCount
		for _, transition := range transitions {
			if removed := transition.Old.Count(); removed > resultingEntryCount {
				return nil, nil, false, errors.New("transition requires removing more entries than exist")
			} else {
				resultingEntryCount -= removed
			}
			resultingEntryCount += transition.New.Count()
		}

		// If the resulting entry count would be too high, then abort the
		// transitioning operation, but return the error as a problem, not an
		// error, since nobody is malfunctioning here.
		if e.maximumEntryCount < resultingEntryCount {
			results := make([]*core.Entry, len(transitions))
			for t, transition := range transitions {
				results[t] = transition.Old
			}
			problems := []*core.Problem{{Error: "transitioning would exceeded allowed entry count"}}
			return results, problems, false, nil
		}
	}

	// Perform the transition.
	results, problems, stagerMissingFiles := core.Transition(
		ctx,
		e.root,
		transitions,
		e.cache,
		e.symbolicLinkMode,
		e.defaultFileMode,
		e.defaultDirectoryMode,
		e.defaultOwnership,
		e.snapshot.DecomposesUnicode,
		e.stager,
	)

	// Determine whether or not the transition made any changes on disk.
	var transitionMadeChanges bool
	for r, result := range results {
		if !result.Equal(transitions[r].Old, true) {
			transitionMadeChanges = true
			break
		}
	}

	// If we're using recursive watching and we made any changes to disk, then
	// send a signal to trigger watch establishment (if needed), because if no
	// watch is currently established due to the synchronization root not having
	// existed, then there's a high likelihood that we just created it.
	if e.watchMode == reifiedWatchModeRecursive && transitionMadeChanges {
		select {
		case e.recursiveWatchRetryEstablish <- struct{}{}:
		default:
		}
	}

	// Ensure that accelerated scanning doesn't return a stale (pre-transition)
	// snapshot. This is critical, especially in the case of poll-based watching
	// (where it has a high chance of occurring). If a pre-transition snapshot
	// is returned by the next call to Scan, then the controller will perform an
	// inversion (on the opposite endpoint) of the transitions that were just
	// applied here. In the case of recursive watching, we just need to ensure
	// that any modified paths get put into the re-check path list, because
	// there could be a delay in the OS reporting the modifications. In the case
	// of poll-based watching, we just need to disable accelerated scanning,
	// which will be automatically re-enabled on the next polling operation. If
	// filesystem watching is disabled, then so is acceleration, and thus
	// there's no way that a stale scan could be returned.
	if e.accelerate {
		if e.watchMode == reifiedWatchModePoll {
			e.accelerate = false
		} else if e.watchMode == reifiedWatchModeRecursive {
			for _, transition := range transitions {
				e.recheckPaths[transition.Path] = true
				if len(e.recheckPaths) > recheckPathsMaximumCapacity {
					e.accelerate = false
					e.recheckPaths = nil
					break
				}
			}
		}
	}

	// If we're using poll-based watching, then strobe the polling channel if
	// Transition made any changes on disk. This is necessary to work around
	// cases where some other mechanism rapidly (and fully) inverts changes, in
	// which case the pre-Transition and post-Transition scans will look the
	// same to the poll-based watching Goroutine and the inversion operation
	// (which should be reported back to the controller) won't be caught. This
	// is unrelated to the stale scan inversion issue mentioned above - in this
	// case the problem is that the changes are seen, but no polling event is
	// ever generated because the polling Goroutine doesn't know what the
	// controller expects the disk to look like - it just knows that nothing has
	// changed between now and some previous point in time.
	//
	// An example of this is when a new file is propagated but then removed by
	// the user before the next poll-based scan. In this case, the polling scan
	// looks the same before and after Transition, and no polling event will be
	// generated if we don't do it here. It's important that we only do this if
	// on-disk changes were actually applied, otherwise we'll drive a feedback
	// loop when problems are encountered for changes that can never be fully
	// applied.
	if e.watchMode == reifiedWatchModePoll && transitionMadeChanges {
		e.strobePollEvents()
	}

	// Wipe the staging directory. We don't monitor for errors here, because we
	// need to return the results and problems no matter what, but if there's
	// something weird going on with the filesystem, we'll see it the next time
	// we scan or stage.
	//
	// TODO: If we see a large number of problems, should we avoid wiping the
	// staging directory? It could be due to an easily correctable error, at
	// which point you wouldn't want to restage if you're talking about lots of
	// files.
	e.stager.wipe()

	// Done.
	return results, problems, stagerMissingFiles, nil
}

// Shutdown implements the Shutdown method for local endpoints.
func (e *endpoint) Shutdown() error {
	// Signal background worker Goroutines to terminate.
	e.workerCancel()

	// Wait for background worker Goroutines to terminate.
	<-e.saveCacheDone
	<-e.watchDone

	// Terminate stager resources.
	e.stager.shutdown()

	// Done.
	return nil
}
