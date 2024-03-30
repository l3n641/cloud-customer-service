package middlewares

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type chatSupportMiddlewareService struct{}

var chatSupportMiddleware = chatSupportMiddlewareService{}

func ChatSupportMiddleware() *clientMiddlewareService {
	return &clientMiddleware
}

func (s *chatSupportMiddlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *chatSupportMiddlewareService) Auth(r *ghttp.Request) {
	ChatSupportAuth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
