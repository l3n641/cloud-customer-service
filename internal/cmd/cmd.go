package cmd

import (
	"cloudCustomerService/internal/controller/admin"
	"cloudCustomerService/internal/controller/client"
	"cloudCustomerService/internal/middlewares"
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
					middlewares.AdminMiddleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)

				group.Bind(
					admin.NewSession().Login,
				)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						middlewares.AdminMiddleware().Auth,
					)
					group.Bind(
						admin.NewAccount(),
						admin.NewMerchant(),
						admin.NewChatSupport(),
					)
				})
			})

			s.Group("/client", func(group *ghttp.RouterGroup) {
				group.Middleware(
					middlewares.ClientMiddleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)

				group.Bind(
					client.NewSession().Login,
				)

			})

			s.Run()
			return nil
		},
	}
)
