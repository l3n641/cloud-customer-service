package admin

import (
	"cloudCustomerService/api/admin/messageTenplate"
	"cloudCustomerService/internal/service"
	"context"
)

func (c *ControllerMessageTenplate) DeleteMessageTemplate(ctx context.Context, req *messageTenplate.DeleteMessageTemplateReq) (res *messageTenplate.DeleteMessageTemplateRes, err error) {
	_, err = service.MessageTemplate().DeleteMessage(ctx, req.Id)
	return &messageTenplate.DeleteMessageTemplateRes{}, err
}
