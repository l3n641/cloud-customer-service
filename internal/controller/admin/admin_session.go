package admin

import (
	"cloudCustomerService/api/admin/session"
	"cloudCustomerService/internal/middlewares"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerSession) Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error) {
	request := g.RequestFromCtx(ctx)
	token, expire := middlewares.Auth().LoginHandler(request.Context())

	return &session.LoginRes{
		Token:      token,
		ExpireTime: expire,
	}, nil
}
