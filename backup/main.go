package main

import (
	_ "login-demo/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"login-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
