package create

import (
	"fmt"
	"path"
	"runtime"

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
	_, currentFilePath, _, _ := runtime.Caller(0)
	fmt.Println(path.Dir(currentFilePath))
	return nil
}
