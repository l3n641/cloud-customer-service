package admin

import (
	"cloudCustomerService/api/admin/account"
	"cloudCustomerService/internal/service"
	"context"
)

func (c *ControllerAccount) UserInfo(ctx context.Context, req *account.UserInfoReq) (res *account.UserInfoRes, err error) {
	adminId := service.BizCtx().GetContext(ctx).GetUserId()
	user, _ := service.Admin().GetUserInfo(ctx, adminId)

	return &account.UserInfoRes{Email: user.Email, NickName: user.Nickname, Avatar: user.Avatar}, nil

}
