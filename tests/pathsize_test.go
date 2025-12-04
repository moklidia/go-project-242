package code

import (
	"code/pathsize"
	"fmt"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSizeFile(t *testing.T) {
	path := filepath.Join("../testdata/testfile.txt")
	human := false
	result, err := pathsize.GetPathSize(path, false, human, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", strconv.Itoa(10), path)
	require.Equal(t, result, expectedResult)
	if result != expectedResult {
		t.Errorf("Actual result %s does not match the expected result %s", result, expectedResult)
	}
}

func TestGetPathSizeFileHumanInBytes(t *testing.T) {
	path := filepath.Join("../testdata/testfile.txt")
	human := true
	result, err := pathsize.GetPathSize(path, false, human, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", "10.00B", path)
	require.Equal(t, result, expectedResult)
	if result != expectedResult {
		t.Errorf("Actual result %s does not match the expected result %s", result, expectedResult)
	}
}

func TestGetPathSizeFileHumanInKiloBytes(t *testing.T) {
	path := filepath.Join("../testdata/image.jpg")
	human := true
	result, err := pathsize.GetPathSize(path, false, human, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", "42.14KB", path)
	require.Equal(t, result, expectedResult)
	if result != expectedResult {
		t.Errorf("Actual result %s does not match the expected result %s", result, expectedResult)
	}
}

func TestGetPathSizeDir(t *testing.T) {
	path := filepath.Join("../testdata/testdir")
	human := false
	result, err := pathsize.GetPathSize(path, false, human, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", strconv.Itoa(46), path)
	require.Equal(t, result, expectedResult)
	if result != expectedResult {
		t.Errorf("Actual result %s does not match the expected result %s", result, expectedResult)
	}
}
