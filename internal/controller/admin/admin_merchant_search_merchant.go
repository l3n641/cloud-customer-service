package admin

import (
	"cloudCustomerService/api"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/merchant"
)

func (c *ControllerMerchant) SearchMerchant(ctx context.Context, req *merchant.SearchMerchantReq) (res *merchant.SearchMerchantRes, err error) {
	out, err := service.Merchant().SearchMerchant(ctx, &model.MerchantSearchInput{
		Page: req.Page,
		Size: req.PageSize,
	})
	return &merchant.SearchMerchantRes{
		PaginationRes: api.PaginationRes{
			Item:     out.List,
			Total:    out.Total,
			Page:     out.Page,
			PageSize: out.Size,
		},
	}, err
}
