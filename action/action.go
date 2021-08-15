package action

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/tal-tech/go-zero/tools/goctl/plugin"
	"github.com/urfave/cli/v2"
)

func Generator(ctx *cli.Context) error {
	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	spew.Dump(p)
	fmt.Printf("plugin args %+v", p)
	return nil
}
