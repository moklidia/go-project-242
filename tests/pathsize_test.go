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
	result, err := pathsize.GetPathSize(path, false, true, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", strconv.Itoa(10), path)
	require.Equal(t, result, expectedResult)
	if result != expectedResult {
		t.Errorf("Actual result %s does not match the expected result %s", result, expectedResult)
	}
}

func TestGetPathSizeDir(t *testing.T) {
	path := filepath.Join("../testdata/testdir")
	result, err := pathsize.GetPathSize(path, false, true, false)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	expectedResult := fmt.Sprintf("%s\t%s", strconv.Itoa(46), path)
	require.Equal(t, result, expectedResult)
	if result != expectedResult {
		t.Errorf("Actual result %s does not match the expected result %s", result, expectedResult)
	}
}
