package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(_ context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			size, err := GetPathSize(path, false, true, false)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(size)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal()
	}
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path)
	if err != nil {
		return "", errors.New(err.Error())
	}

	formattedOutput := fmt.Sprintf("%s\t%s", strconv.Itoa(size), path)

	return formattedOutput, nil
}

func GetSize(path string) (int, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, errors.New("the directory or file does not exist")
	}
	if fileInfo.IsDir() {
		size, err := GetDirSize(path)
		if err != nil {
			return 0, errors.New(err.Error())
		}
		return size, nil
	} else {
		return int(fileInfo.Size()), nil
	}
}

func GetDirSize(path string) (int, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, errors.New("couldn't read the directory")
	}
	size := 0
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return 0, errors.New("couldn't get file size")
		}
		size += int(info.Size())
	}
	return size, nil
}
