package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testFiles = []string{"file1.txt", "file2.mp3", "file3"}

func createTestFiles() ([]*os.File, error) {

	var (
		testFiles    = []string{"file1.txt", "file2.mp3", "file3"}
		tmpDir       = os.TempDir()
		createdFiles []*os.File
	)

	for _, file := range testFiles {
		createdFile, err := os.Create(filepath.Join(tmpDir, file))
		if err != nil {
			removeTestFiles(createdFiles)
			return nil, err
		}

		createdFiles = append(createdFiles, createdFile)
	}

	return createdFiles, nil
}

func removeTestFiles(createdFiles []*os.File) {
	tmpDir := os.TempDir()

	for _, f := range testFiles {
		_ = os.Remove(filepath.Join(tmpDir, f))
	}

	for _, cf := range createdFiles {
		_ = cf.Close()
	}
}

func TestScanForTranscriptFiles(t *testing.T) {

	createdFiles, err := createTestFiles()
	require.NoError(t, err)

	err = ScanForTranscriptFiles(os.TempDir(), func(file *os.DirEntry) {
		assert.Equal(t, filepath.Ext((*file).Name()), SupportedFileExtensions)
	})

	require.NoError(t, err)

	removeTestFiles(createdFiles)
}
