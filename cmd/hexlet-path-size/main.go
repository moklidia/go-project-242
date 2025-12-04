package main

import (
	"code/pathsize"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(_ context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			size, err := pathsize.GetPathSize(path, false, true, false)
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
