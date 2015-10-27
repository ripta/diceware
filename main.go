package main

// import "fmt"
import "os"

import "github.com/codegangsta/cli"

var genLength int
var genShort bool
var genQuiet bool

var version = "1.0.0"

func main() {
	app := cli.NewApp()
	
	app.Name = "diceware"
	app.Version = version

	bannerBytes, _ := dataBannerBytes()
	app.Usage = string(bannerBytes)

	app.Action = GeneratePassphrase
	app.Flags = []cli.Flag {
		cli.IntFlag{
			Name:  "n, num-words",
			Value: 6,
			Usage: "number of words in the passphrase",
		},
		cli.BoolFlag{
			Name:  "q, quiet",
			Usage: "when specified, only prints words, and not the generated numbers",
		},
		cli.BoolFlag{
			Name:  "s, short",
			Usage: "when specified, prints short form: no spaces between words",
		},
		cli.StringFlag{
			Name:  "w, word-list",
			Value: "diceware",
			Usage: "name of word list to use (see `diceware list` output)",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l", "li", "lis"},
			Usage:   "list all available wordlists",
			Action:  ListWordlistNames,
		},
	}

	app.Run(os.Args)
}
