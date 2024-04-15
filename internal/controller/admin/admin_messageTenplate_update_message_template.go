package admin

import (
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/messageTenplate"
)

func (c *ControllerMessageTenplate) UpdateMessageTemplate(ctx context.Context, req *messageTenplate.UpdateMessageTemplateReq) (res *messageTenplate.UpdateMessageTemplateRes, err error) {
	_, err = service.MessageTemplate().UpdateMessageTemplate(ctx, req.Id, req.MsgType, req.Lang, req.Content)
	return &messageTenplate.UpdateMessageTemplateRes{}, err
}
