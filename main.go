package main

import (
	"cloudCustomerService/internal/cmd"
	_ "cloudCustomerService/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"

	_ "cloudCustomerService/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
