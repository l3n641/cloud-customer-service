package chatSupport

import (
	"cloudCustomerService/internal/middlewares"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"cloudCustomerService/api/chatSupport/session"
)

func (c *ControllerSession) Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error) {
	request := g.RequestFromCtx(ctx)
	token, expire := middlewares.ChatSupportAuth().LoginHandler(request.Context())

	return &session.LoginRes{
		Token:      token,
		ExpireTime: expire,
	}, nil
}
