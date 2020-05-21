package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "what-is-this-secret"
	app.Version = version
	app.EnableBashCompletion = true
	app.Authors = []*cli.Author{{Name: "Madhu Akula"}}
	app.Usage = "An utility to check what secret you have!"

	app.Before = func(context *cli.Context) error {
		if context.Bool("verbose") {
			log.SetFlags(0)
		} else {
			log.SetOutput(ioutil.Discard)
		}
		return nil
	}

	app.Compiled = time.Now()

	app.Commands = []*cli.Command{
		{
			Name: "search",
			// Aliases: []string{""},
			Usage: "Search the secret key",

			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "key",
					Aliases:  []string{"k"},
					Usage:    "Provide the secret key to check",
					Required: true,
				},
			},
			// Action: func(c *cli.Context) error {
			// 	secretCheck(c.String("key"))
			// 	return nil
			// },
			Action: cliSecretCheck,
		},
		{
			Name: "server",
			// Aliases: []string{""},
			Usage:  "Run the what-is-this-secret in server mode",
			Action: runServer,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		// Run with '--help' for usage.
		log.Fatal(err)
	}
}
