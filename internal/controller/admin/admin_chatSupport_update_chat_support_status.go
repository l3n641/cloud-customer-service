package admin

import (
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/chatSupport"
)

func (c *ControllerChatSupport) UpdateChatSupportStatus(ctx context.Context, req *chatSupport.UpdateChatSupportStatusReq) (res *chatSupport.UpdateChatSupportStatusRes, err error) {
	_, err = service.ChatSupport().UpdateChatSupportStatus(ctx, req.Id, req.Status)
	return &chatSupport.UpdateChatSupportStatusRes{}, err
}
