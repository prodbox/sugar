package create

import (
	"fmt"

	"github.com/prodbox/sugar/os/upath"
	"github.com/urfave/cli/v2"
)

var (
	DefaultCommand = &cli.Command{
		Name:   "create",
		Usage:  "创建新项目",
		Before: Before,
		Action: Action,
	}
)

func Before(ctx *cli.Context) error {
	return nil
}

func Action(ctx *cli.Context) error {
	fmt.Println(upath.ExecutableFolder())
	return nil
}
