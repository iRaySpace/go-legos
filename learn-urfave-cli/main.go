package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var host string
	var port int
	var live bool

	app := &cli.App{
		Name:  "learn-urfave-cli",
		Usage: "learning urfave/cli",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "host",
				Aliases:     []string{"H"},
				Usage:       "HOST of the server",
				Destination: &host,
			},
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Value:       8000,
				Destination: &port,
				Action: func(c *cli.Context, v int) error {
					if v > 65535 || v < 0 {
						return fmt.Errorf("out of range[0-65535] port: %d", v)
					}
					return nil
				},
			},
			&cli.BoolFlag{
				Name:        "live",
				Aliases:     []string{"l"},
				Usage:       "live mode",
				Destination: &live,
			},
		},
		Action: func(c *cli.Context) error {
			if host != "localhost" {
				fmt.Println("Why not localhost?!?")
				return nil
			}
			if c.Bool("live") {
				fmt.Println("Running in live mode...")
			}
			fmt.Println("Running on localhost:", port)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
