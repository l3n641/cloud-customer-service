package chatSupport

import (
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/chatSupport/account"
)

func (c *ControllerAccount) UserInfo(ctx context.Context, req *account.UserInfoReq) (res *account.UserInfoRes, err error) {
	request := g.RequestFromCtx(ctx)
	userId := gconv.Int(middlewares.ChatSupportAuth().GetIdentity(request.Context()))
	user, err := service.ChatSupport().GetUserInfo(ctx, userId)

	return &account.UserInfoRes{Email: user.Email, NickName: user.Nickname, Avatar: user.Avatar}, err
}
