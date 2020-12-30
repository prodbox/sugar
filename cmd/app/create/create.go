package create

import (
	"fmt"
	"github.com/prodbox/sugar/os/upath"
	"path"

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
	serverPath := path.Join(upath.ExecutableSourceRoot(), "server")
	fmt.Println(serverPath)
	if upath.IsExist(serverPath) == false {
		if err := upath.Mkdir(serverPath); err != nil {
			return err
		}
	}
	return nil
}
