package admin

import (
	"cloudCustomerService/api/admin/messageTenplate"
	"cloudCustomerService/internal/service"
	"context"
)

func (c *ControllerMessageTenplate) CreateMessageTemplate(ctx context.Context, req *messageTenplate.CreateMessageTemplateReq) (res *messageTenplate.CreateMessageTemplateRes, err error) {

	id, err := service.MessageTemplate().CreateMessageTemplate(ctx, req.MsgType, req.Lang, req.Content)
	return &messageTenplate.CreateMessageTemplateRes{Id: id}, err
}
