package admin

import (
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/chatSupport"
)

func (c *ControllerChatSupport) UpdateChatSupport(ctx context.Context, req *chatSupport.UpdateChatSupportReq) (res *chatSupport.UpdateChatSupportRes, err error) {
	err = service.ChatSupport().UpdateChatSupport(ctx, &model.ChatSupportUpdateInput{
		Id:            req.Id,
		NickName:      req.Nickname,
		Password:      req.Password,
		MerchantGroup: req.MerchantGroup,
	},
	)

	return &chatSupport.UpdateChatSupportRes{}, err
}
