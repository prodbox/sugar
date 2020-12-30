package app

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

var (
	defaultCmd = cli.NewApp()
)

func init() {
	defaultCmd.Name = "Sugar"
	defaultCmd.Usage = "Golang业务框架"
	defaultCmd.HideVersion = true

	cli.HelpFlag = &cli.BoolFlag{
		Name:  "help,h",
		Usage: "查看帮助信息",
	}
}

// Register 注册
func Register(cmds ...*cli.Command) {
	defaultCmd.Commands = append(defaultCmd.Commands, cmds...)

	// 排序
	sort.Slice(defaultCmd.Commands, func(i, j int) bool {
		return defaultCmd.Commands[i].Name < defaultCmd.Commands[j].Name
	})
}

// Execute 运行
func Execute() error {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			os.Exit(1)
		}
	}()
	return defaultCmd.Run(os.Args)
}
