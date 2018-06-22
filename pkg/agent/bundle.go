package agent

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"

	"github.com/havoc-io/mutagen/pkg/process"
)

const (
	agentBundleName = "mutagen-agents.tar.gz"
)

func executableForPlatform(goos, goarch string) (string, error) {
	// Compute the path to the current executable.
	executablePath, err := os.Executable()
	if err != nil {
		return "", errors.Wrap(err, "unable to determine executable path")
	}

	// Compute the path to the agent bundle.
	bundlePath := filepath.Join(filepath.Dir(executablePath), agentBundleName)

	// Open the bundle path and ensure its closure.
	bundle, err := os.Open(bundlePath)
	if err != nil {
		return "", errors.Wrap(err, "unable to open agent bundle")
	}
	defer bundle.Close()

	// Create a decompressor and ensure its closure.
	bundleDecompressor, err := gzip.NewReader(bundle)
	if err != nil {
		return "", errors.Wrap(err, "unable to decompress agent bundle")
	}
	defer bundleDecompressor.Close()

	// Create an archive reader.
	bundleArchive := tar.NewReader(bundleDecompressor)

	// Scan until we find a matching header.
	var header *tar.Header
	for {
		if h, err := bundleArchive.Next(); err != nil {
			if err == io.EOF {
				break
			}
			return "", errors.Wrap(err, "unable to read archive header")
		} else if h.Name == fmt.Sprintf("%s_%s", goos, goarch) {
			header = h
			break
		}
	}

	// Check if we have a valid header. If not, there was no match.
	if header == nil {
		return "", errors.New("unsupported platform")
	}

	// Compute the base name for the output file.
	targetBaseName := process.ExecutableName(agentBaseName, goos)

	// Create a temporary file in which to receive the agent on disk.
	file, err := ioutil.TempFile("", targetBaseName)
	if err != nil {
		return "", errors.Wrap(err, "unable to create temporary file")
	}

	// Copy data into the file.
	if _, err := io.CopyN(file, bundleArchive, header.Size); err != nil {
		file.Close()
		os.Remove(file.Name())
		return "", errors.Wrap(err, "unable to copy agent data")
	}

	// If we're not on Windows and our target system is not Windows, mark the
	// file as executable. This will save us an additional "chmod +x" command
	// during agent installation.
	if runtime.GOOS != "windows" && goos != "windows" {
		if err := file.Chmod(0700); err != nil {
			file.Close()
			os.Remove(file.Name())
			return "", errors.Wrap(err, "unable to make agent executable")
		}
	}

	// Close the file.
	if err := file.Close(); err != nil {
		os.Remove(file.Name())
		return "", errors.Wrap(err, "unable to close temporary file")
	}

	// Success.
	return file.Name(), nil
}

// CopyBundleForTesting copies the agent bundle from it's build path alongside
// the current executable. It is useful for copying the agent bundle next to the
// current test executable.
func CopyBundleForTesting() error {
	// Compute the path to the test executable and its parent directory.
	executablePath, err := os.Executable()
	if err != nil {
		return errors.Wrap(err, "unable to compute test executable path")
	}
	testDirectory := filepath.Dir(executablePath)

	// Compute the path to the agent bundle in the $GOPATH/bin directory.
	var agentBundlePath string
	if gopath := os.Getenv("GOPATH"); gopath == "" {
		return errors.New("GOPATH not set")
	} else {
		agentBundlePath = filepath.Join(gopath, "bin", agentBundleName)
	}

	// Create a file that will be a copy of the agent bundle.
	// HACK: We're assuming that Go runs test executables inside temporary
	// directories that it cleans up, which does seem to be the case, but it'd
	// be nice if there were some way to remove the agent bundle ourselves,
	// maybe with some sort of atexit-like function.
	bundleCopyFile, err := os.Create(filepath.Join(testDirectory, agentBundleName))
	if err != nil {
		return errors.Wrap(err, "unable to create agent bundle copy file")
	}
	defer bundleCopyFile.Close()

	// Open the agent bundle.
	bundleFile, err := os.Open(agentBundlePath)
	if err != nil {
		return errors.Wrap(err, "unable to open agent bundle file")
	}
	defer bundleFile.Close()

	// Copy agent bundle contents.
	if _, err := io.Copy(bundleCopyFile, bundleFile); err != nil {
		return errors.Wrap(err, "unable to copy bundle file contents")
	}

	// Success.
	return nil
}
