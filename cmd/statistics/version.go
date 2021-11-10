package main

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli"
)

var (
	// current project version
	version = ""
	// current git commit hash
	commit = ""
	// compile date
	date = ""
	// GoVersion system go version
	GoVersion = runtime.Version()
	// Platform info
	Platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
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
	version := fmt.Sprintf("Bridge version: %s-%s\n", version, commit)
	if all {
		version += fmt.Sprintf("App build date: %s\n", date)
		version += fmt.Sprintf("System version: %s\n", Platform)
		version += fmt.Sprintf("Golang version: %s\n", GoVersion)
	}

	return version
}
