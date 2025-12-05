package code

import (
	"code"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSizeFile(t *testing.T) {
	path := filepath.Join("../testdata/testfile.txt")
	human := false
	result, err := code.GetPathSize(path, false, human, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	require.Equal(t, "10B", result)
}

func TestGetPathSizeFileHumanInBytes(t *testing.T) {
	path := filepath.Join("../testdata/testfile.txt")
	human := true
	result, err := code.GetPathSize(path, false, human, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	require.Equal(t, "10B", result)
}

func TestGetPathSizeFileHumanInKiloBytes(t *testing.T) {
	path := filepath.Join("../testdata/image.jpg")
	human := true
	result, err := code.GetPathSize(path, false, human, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	require.Equal(t, "42.14KB", result)
}

func TestGetPathSizeDirAll(t *testing.T) {
	path := filepath.Join("../testdata/testdir")
	human := false
	all := false
	result, err := code.GetPathSize(path, false, human, all)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	require.Equal(t, "46B", result)
}

func TestGetPathSizeDirWithoutHidden(t *testing.T) {
	path := filepath.Join("../testdata/testdir")
	human := false
	all := true
	result, err := code.GetPathSize(path, false, human, all)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	require.Equal(t, "66B", result)
}

func TestGetPathSizeDirRecursive(t *testing.T) {
	path := filepath.Join("../testdata/recursive")
	recursive := true
	result, err := code.GetPathSize(path, recursive, false, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	require.Equal(t, "84B", result)
}

func TestGetPathSizeDirNonRecursive(t *testing.T) {
	path := filepath.Join("../testdata/recursive")
	recursive := false
	result, err := code.GetPathSize(path, recursive, false, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	require.Equal(t, "15B", result)
}
