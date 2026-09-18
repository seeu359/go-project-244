package main

import (
	"context"
	"fmt"
	"os"

	"code"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Usage:   "output format",
				Value:   "stylish",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			filePath1 := cmd.Args().Get(0)
			filePath2 := cmd.Args().Get(1)
			format := cmd.String("format")

			diff, err := code.GenDiff(filePath1, filePath2, format)
			if err != nil {
				return err
			}

			fmt.Println(diff)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
