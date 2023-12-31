package cmd

import (
	"cloudCustomerService/internal/controller/admin"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/admin", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.AdminAuthMiddleware().Ctx,
					ghttp.MiddlewareHandlerResponse)

				group.Bind(
					admin.NewSession().Login,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.AdminAuthMiddleware().Auth)
					group.Bind(
						admin.NewAccount(),
						admin.NewCustomerServiceAgent(),
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
