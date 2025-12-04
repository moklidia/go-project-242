package code

import (
	"code"
	"fmt"
	"path/filepath"
	"strconv"
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

	expectedResult := fmt.Sprintf("%s\t%s", strconv.Itoa(10), path)
	require.Equal(t, expectedResult, result)
}

func TestGetPathSizeFileHumanInBytes(t *testing.T) {
	path := filepath.Join("../testdata/testfile.txt")
	human := true
	result, err := code.GetPathSize(path, false, human, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", "10.00B", path)
	require.Equal(t, expectedResult, result)
}

func TestGetPathSizeFileHumanInKiloBytes(t *testing.T) {
	path := filepath.Join("../testdata/image.jpg")
	human := true
	result, err := code.GetPathSize(path, false, human, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", "42.14KB", path)
	require.Equal(t, expectedResult, result)
}

func TestGetPathSizeDirAll(t *testing.T) {
	path := filepath.Join("../testdata/testdir")
	human := false
	all := false
	result, err := code.GetPathSize(path, false, human, all)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", strconv.Itoa(46), path)
	require.Equal(t, expectedResult, result)
}

func TestGetPathSizeDirWithoutHidden(t *testing.T) {
	path := filepath.Join("../testdata/testdir")
	human := false
	all := true
	result, err := code.GetPathSize(path, false, human, all)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", strconv.Itoa(66), path)
	require.Equal(t, expectedResult, result)
}

func TestGetPathSizeDirRecursive(t *testing.T) {
	path := filepath.Join("../testdata/recursive")
	recursive := true
	result, err := code.GetPathSize(path, recursive, false, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", strconv.Itoa(84), path)
	require.Equal(t, expectedResult, result)
}

func TestGetPathSizeDirNonRecursive(t *testing.T) {
	path := filepath.Join("../testdata/recursive")
	recursive := false
	result, err := code.GetPathSize(path, recursive, false, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", strconv.Itoa(15), path)
	require.Equal(t, expectedResult, result)
}
