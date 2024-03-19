package admin

import (
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/admin/account"
)

func (c *ControllerAccount) UserInfo(ctx context.Context, req *account.UserInfoReq) (res *account.UserInfoRes, err error) {
	request := g.RequestFromCtx(ctx)
	adminId := gconv.Int(middlewares.Auth().GetIdentity(request.Context()))
	user, _ := service.Admin().GetUserInfo(ctx, adminId)

	return &account.UserInfoRes{Username: user.Username, NickName: user.Nickname, Avatar: user.Avatar}, nil
}
