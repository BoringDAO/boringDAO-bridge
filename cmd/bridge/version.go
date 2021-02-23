package main

import (
	"fmt"

	"github.com/boringdao/bridge"
	"github.com/urfave/cli"
)

var versionCMD = cli.Command{
	Name:  "version",
	Usage: "Show version about app",
	Action: func(ctx *cli.Context) error {
		fmt.Print(getVersion(true))

		return nil
	},
}

func getVersion(all bool) string {
	version := fmt.Sprintf("bridge-node version: %s-%s\n", bridge.CurrentVersion, bridge.CurrentCommit)
	if all {
		version += fmt.Sprintf("App build date: %s\n", bridge.BuildDate)
		version += fmt.Sprintf("System version: %s\n", bridge.Platform)
		version += fmt.Sprintf("Golang version: %s\n", bridge.GoVersion)
	}

	return version
}
