package middlewares

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type clientMiddlewareService struct{}

var clientMiddleware = clientMiddlewareService{}

func ClientMiddleware() *clientMiddlewareService {
	return &clientMiddleware
}

func (s *clientMiddlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *clientMiddlewareService) Auth(r *ghttp.Request) {
	action := r.Get("action")
	if action.IsEmpty() || action.String() != "login" {
		ClientAuth().MiddlewareFunc()(r)
	}
	r.Middleware.Next()
}
