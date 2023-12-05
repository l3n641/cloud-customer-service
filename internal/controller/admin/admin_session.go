package admin

import (
	"cloudCustomerService/api/admin/session"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerSession) Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error) {
	request := g.RequestFromCtx(ctx)

	loginOutput, err := service.Admin().Login(ctx, &model.AdminLoginInput{
		Email:    req.Email,
		Password: req.Password,
		Ip:       request.GetRemoteIp(),
	},
	)
	if err != nil {
		return nil, err

	}

	service.BizCtx().SetData(ctx, loginOutput.Data)

	auth := service.AdminAuthMiddleware().GetJWTMiddleware(ctx)
	token, expireTime := auth.LoginHandler(ctx)
	return &session.LoginRes{
		Token:      token,
		ExpireTime: expireTime,
	}, nil
}
