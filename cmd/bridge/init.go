package main

import (
	"bufio"
	"fmt"
	"os"

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
		},
	}
}

func initialize(ctx *cli.Context) error {
	repoRoot, err := repo.PathRootWithDefault(ctx.GlobalString("repo"))
	if err != nil {
		return err
	}

	fmt.Printf("initializing bridge at %s\n", repoRoot)

	if repo.Initialized(repoRoot) {
		fmt.Println("bridge configuration file already exists")
		fmt.Println("reinitializing would overwrite your configuration, Y/N?")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text() == "Y" || input.Text() == "y" {
			return repo.Initialize(repoRoot)
		}
		return nil
	}

	return repo.Initialize(repoRoot)
}
