package synchronization

import (
	"fmt"
	"path/filepath"

	"github.com/mutagen-io/mutagen/pkg/filesystem"
)

// pathForSession computes the path to the serialized session for the given
// session identifier. An empty session identifier will return the sessions
// directory path.
func pathForSession(dataDir filesystem.DataDirFunc, sessionIdentifier string) (string, error) {
	// Compute/create the sessions directory.
	sessionsDirectoryPath, err := dataDir(true, filesystem.MutagenSynchronizationSessionsDirectoryName)
	if err != nil {
		return "", fmt.Errorf("unable to compute/create sessions directory: %w", err)
	}

	// Success.
	return filepath.Join(sessionsDirectoryPath, sessionIdentifier), nil
}

// pathForArchive computes the path to the serialized archive for the given
// session identifier.
func pathForArchive(dataDir filesystem.DataDirFunc, session string) (string, error) {
	// Compute/create the archives directory.
	archivesDirectoryPath, err := dataDir(true, filesystem.MutagenSynchronizationArchivesDirectoryName)
	if err != nil {
		return "", fmt.Errorf("unable to compute/create archives directory: %w", err)
	}

	// Success.
	return filepath.Join(archivesDirectoryPath, session), nil
}
