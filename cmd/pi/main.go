package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nanlei2000/pi/pi"
	"github.com/urfave/cli/v2"
)

func main() {
	var n uint64
	var print bool

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
			&cli.BoolFlag{
				Name:        "print",
				Aliases:     []string{"p"},
				Value:       true,
				Usage:       "print pi or not",
				Destination: &print,
			},
		},
		Action: func(c *cli.Context) error {
			start := time.Now()
			piRes := pi.Calc(n)
			end := time.Now()
			fmt.Printf("time: %v \n", end.Sub(start))
			if print {
				fmt.Println(pi.Fmt(piRes))
			}
			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
