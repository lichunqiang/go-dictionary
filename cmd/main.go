package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/lichunqiang/dictionary/actions"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "Dictionary"
	app.Usage = "dictionary for everyone!"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		{
			Name:  "lichunqiang",
			Email: "light-li@hotmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "en",
			Usage: "Language for translate",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("lang") == "en" {
			fmt.Println("tranlate to english")
		} else {
			cli.ShowAppHelp(c)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "hi",
			Aliases: []string{"h"},
			Usage:   "This is for command",
			Action:  actions.ActionTest,
		},
	}

	app.CommandNotFound = func(c *cli.Context, ex string) {
		fmt.Println(ex)
	}

	app.Run(os.Args)
}
