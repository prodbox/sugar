package cmd

import (
	"sugar/cmd/app"
	"sugar/cmd/app/create"
)

func init() {
	// 创建项目
	app.Register(create.DefaultCommand)
}

func Execute() {
	app.Execute()
}
