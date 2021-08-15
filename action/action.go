package action

import (
	"fmt"
	"github.com/tal-tech/go-zero/tools/goctl/plugin"
	"github.com/urfave/cli/v2"
)

func Generator(ctx *cli.Context) error {
	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	fmt.Printf("plugin args %+v", p)
	return nil
}
