package admin

import (
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/chatSupport"
)

func (c *ControllerChatSupport) DetailChatSupport(ctx context.Context, req *chatSupport.DetailChatSupportReq) (res *chatSupport.DetailChatSupportRes, err error) {
	data, err := service.ChatSupport().GetDetailById(ctx, req.Id)
	return &chatSupport.DetailChatSupportRes{
		Id:            data.Id,
		Email:         data.Email,
		MerchantGroup: data.MerchantGroup,
	}, err
}
