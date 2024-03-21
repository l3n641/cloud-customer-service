package admin

import (
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/chatSupport"
)

func (c *ControllerChatSupport) CreateChatSupport(ctx context.Context, req *chatSupport.CreateChatSupportReq) (res *chatSupport.CreateChatSupportRes, err error) {
	err = service.ChatSupport().CreteChatSupport(ctx, &model.ChatSupportCreateInput{
		Email:         req.Email,
		Password:      req.Password,
		NickName:      req.Nickname,
		MerchantGroup: req.MerchantGroup,
	})
	if err == nil {
		return &chatSupport.CreateChatSupportRes{}, nil
	}
	return nil, err
}
