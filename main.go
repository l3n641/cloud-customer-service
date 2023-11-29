package main

import (
	_ "cloudCustomerService/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"cloudCustomerService/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
