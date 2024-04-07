package main

import (
	_ "cloudCustomerService/internal/logic"

	"cloudCustomerService/internal/cmd"
	_ "cloudCustomerService/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
