package admin

import (
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/messageTenplate"
)

func (c *ControllerMessageTenplate) GetMessageTemplateDetail(ctx context.Context, req *messageTenplate.GetMessageTemplateDetailReq) (res *messageTenplate.GetMessageTemplateDetailRes, err error) {

	data, err := service.MessageTemplate().GetMessageDetail(ctx, req.Id)
	return &messageTenplate.GetMessageTemplateDetailRes{
		Id:      data.Id,
		MsgType: data.MsgType,
		Lang:    data.Lang,
		Content: data.Content,
	}, err
}
