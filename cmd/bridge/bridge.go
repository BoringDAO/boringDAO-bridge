package main

import (
	"os"
	"time"

	"github.com/boringdao/bridge/pkg/kit/log"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var logger = log.NewWithModule("cmd")

func main() {
	app := cli.NewApp()
	app.Name = "Boring"
	app.Usage = "Boring Node"
	app.Compiled = time.Now()

	// global flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "repo",
			Usage: "Boring repository path",
		},
	}

	app.Commands = []cli.Command{
		startCMD,
		versionCMD,
		initCMD(),
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Red(err.Error())
		os.Exit(-1)
	}
}
