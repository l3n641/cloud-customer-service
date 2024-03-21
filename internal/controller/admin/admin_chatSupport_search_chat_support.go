package admin

import (
	"cloudCustomerService/api"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/chatSupport"
)

func (c *ControllerChatSupport) SearchChatSupport(ctx context.Context, req *chatSupport.SearchChatSupportReq) (res *chatSupport.SearchChatSupportRes, err error) {
	out, err := service.ChatSupport().SearchChatSupport(ctx, &model.ChatSupportSearchInput{
		Page:    req.Page,
		Size:    req.PageSize,
		OrderBy: req.OrderReq.String(),
		SearchChatSupportFields: model.SearchChatSupportFields{
			Email:    req.Email,
			IsOnline: req.IsOnline,
			Status:   req.Status,
		},
	})
	
	return &chatSupport.SearchChatSupportRes{
		PaginationRes: api.PaginationRes{
			Item:     out.List,
			Total:    out.Total,
			Page:     out.Page,
			PageSize: out.Size,
		},
	}, err
}
