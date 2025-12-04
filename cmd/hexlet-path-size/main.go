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
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "human",
				Value: false,
				Usage: "human-readable sizes (auto-select unit)",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			path := "."
			if cmd.NArg() > 0 {
				path = cmd.Args().Get(0)
			}
			human := cmd.Bool("human")
			size, err := pathsize.GetPathSize(path, false, human, false)
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
