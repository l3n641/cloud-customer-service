package admin

import (
	"cloudCustomerService/api/admin/session"
	"cloudCustomerService/internal/middlewares"
	"context"
)

func (c *ControllerSession) SignOut(ctx context.Context, req *session.SignOutReq) (res *session.SignOutRes, err error) {
	middlewares.Auth().LogoutHandler(ctx)
	return
}
