package create

import (
	"fmt"
	"path"

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

	// go get 源码所在目录
	defaultSourcePath = upath.ExecutableSourceRoot()
)

func Before(ctx *cli.Context) error {
	if upath.IsWritable(defaultSourcePath) == false {
		return fmt.Errorf("%s 写权限不足, 请切换至超级用户", defaultSourcePath)
	}
	return nil
}

func Action(ctx *cli.Context) error {
	serverPath := path.Join(defaultSourcePath, "server")
	if upath.IsExist(serverPath) == false {
		if err := upath.Mkdir(serverPath); err != nil {
			return err
		}
	}
	return nil
}
