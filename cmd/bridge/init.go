package main

import (
	"fmt"

	"github.com/boringdao/bridge/internal/repo"
	"github.com/urfave/cli"
)

func initCMD() cli.Command {
	return cli.Command{
		Name:   "init",
		Usage:  "Initialize bridge local configuration",
		Action: initialize,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config",
				Value: "",
				Usage: "bridge config repo path",
			},
			cli.BoolFlag{
				Name:  "interactive,i",
				Usage: "configure bridge in interactive mode",
			},
		},
	}
}

func initialize(ctx *cli.Context) error {
	repoRoot, err := repo.PathRootWithDefault(ctx.GlobalString("repo"))
	if err != nil {
		return err
	}

	interactive := ctx.Bool("interactive")

	fmt.Printf("initializing bridge at %s\n", repoRoot)

	if repo.Initialized(repoRoot) {
		fmt.Println("bridge configuration file already exists")
		fmt.Println("reinitializing would overwrite your configuration, Y/n?")
		if repo.ReadYes() {
			return repo.Initialize(repoRoot, interactive)
		}
		return nil
	}

	return repo.Initialize(repoRoot, interactive)
}
