package client

import (
	"cloudCustomerService/internal/middlewares"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"cloudCustomerService/api/client/session"
)

func (c *ControllerSession) Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error) {
	request := g.RequestFromCtx(ctx)
	token, expire := middlewares.ClientAuth().LoginHandler(request.Context())

	return &session.LoginRes{
		Token:      token,
		ExpireTime: expire,
	}, nil
}
