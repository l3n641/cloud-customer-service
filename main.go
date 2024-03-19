package main

import (
	_ "cloudCustomerService/internal/packed"

	_ "cloudCustomerService/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"cloudCustomerService/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
