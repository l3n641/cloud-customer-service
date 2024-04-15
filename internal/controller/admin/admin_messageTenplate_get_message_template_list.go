package admin

import (
	"cloudCustomerService/api"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/messageTenplate"
)

func (c *ControllerMessageTenplate) GetMessageTemplateList(ctx context.Context, req *messageTenplate.GetMessageTemplateListReq) (res *messageTenplate.GetMessageTemplateListRes, err error) {
	out, err := service.MessageTemplate().GetMessageList(ctx, &model.MessageTemplateListInput{
		MsgType:    req.MsgType,
		Lang:       req.Lang,
		Pagination: model.Pagination{Page: req.Page, Size: req.PageSize},
	})
	return &messageTenplate.GetMessageTemplateListRes{
		PaginationRes: api.PaginationRes{
			Item:     out.List,
			Total:    out.Total,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}, err
}
