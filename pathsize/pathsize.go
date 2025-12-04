package pathsize

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, all, recursive)
	if err != nil {
		return "", errors.New(err.Error())
	}

	formattedOutput := Format(size, path, human)

	return formattedOutput, nil
}

func getSize(path string, all bool, recursive bool) (int, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, errors.New("the directory or file does not exist")
	}
	if fileInfo.IsDir() {
		size, err := getDirSize(path, all, recursive)
		if err != nil {
			return 0, errors.New(err.Error())
		}
		return size, nil
	} else {
		return int(fileInfo.Size()), nil
	}
}

func getDirSize(path string, all bool, recursive bool) (int, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, errors.New("couldn't read the directory")
	}

	size := 0
	for _, file := range files {
		if !all && strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if file.IsDir() {
			if recursive {
				nestedPath := filepath.Join(path, "/", file.Name())
				dirSize, err := getDirSize(nestedPath, all, recursive)
				if err != nil {
					return 0, errors.New("couldn't get file info")
				}
				size += dirSize
			}
		} else {
			info, err := file.Info()
			if err != nil {
				return 0, errors.New("couldn't get file info")
			}
			size += int(info.Size())
		}
	}
	return size, nil
}
