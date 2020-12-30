package cmd

import (
	"github.com/prodbox/sugar/cmd/app"
	"github.com/prodbox/sugar/cmd/app/create"
)

func init() {
	// 创建项目
	app.Register(create.DefaultCommand)
}

func Execute() {
	app.Execute()
}
