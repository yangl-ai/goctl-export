package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yangl-ai/goctl-export/action"
	"os"
	"runtime"
)

var (
	version  = ""
	commands = []*cli.Command{
		{
			Name:   "doc-export",
			Usage:  "generates swagger.json",
			Action: action.Generator,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "type",
					Usage: "api request address",
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a plugin of goctl to generate swagger.json"
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-swagger: %+v\n", err)
	}
}
