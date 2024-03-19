package middlewares

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type adminMiddlewareService struct{}

var adminMiddleware = adminMiddlewareService{}

func AdminMiddleware() *adminMiddlewareService {
	return &adminMiddleware
}

func (s *adminMiddlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *adminMiddlewareService) Auth(r *ghttp.Request) {
	Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
