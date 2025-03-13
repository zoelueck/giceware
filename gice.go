package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type diceroll struct {
	die0 int8
	die1 int8
	die2 int8
	die3 int8
	die4 int8
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:    "words",
				Usage:   "set the amount of words your password should have",
				Aliases: []string{"w"},
				Value:   6,
			},
			&cli.StringFlag{
				Name:    "config",
				Usage:   "Override default config",
				Aliases: []string{"c"},
				Value:   "~/.config/gice/giceConfig.toml",
			},
			&cli.Int64Flag{
				Name:    "numbers",
				Usage:   "set the number or numbers, please note that if the number of numbers is larger than the  number of words you will get multi digit numbers",
				Aliases: []string{"n"},
				Value:   1,
			},
		},
		Action: func(ctx *cli.Context) error {
			i := ctx.Int64("words")
			if i <= 0 {
				return fmt.Errorf("can't set 0 or less words")
			} else {
				fmt.Printf("Words: %v\n", i)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
