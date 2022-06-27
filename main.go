package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var n uint64

	app := &cli.App{
		Name:  "pi",
		Usage: "calculate pi",
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:        "n",
				Value:       100,
				Usage:       "digit",
				Destination: &n,
			},
		},
		Action: func(c *cli.Context) error {
			pi := Calc(n).String()
			fmt.Printf("%s.%s", string(pi[0]), string(pi[1:]))

			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}