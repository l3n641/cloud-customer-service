package chatSupport

import (
	"cloudCustomerService/internal/middlewares"
	"context"

	"cloudCustomerService/api/chatSupport/session"
)

func (c *ControllerSession) SignOut(ctx context.Context, req *session.SignOutReq) (res *session.SignOutRes, err error) {
	middlewares.ChatSupportAuth().LogoutHandler(ctx)
	return
}
