package admin

import (
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/merchant"
)

func (c *ControllerMerchant) DetailMerchant(ctx context.Context, req *merchant.DetailMerchantReq) (res *merchant.DetailMerchantRes, err error) {
	data, err := service.Merchant().GetDetailById(ctx, req.Id)
	return &merchant.DetailMerchantRes{
		Id:           data.Id,
		MerchantId:   data.MerchantId,
		MerchantName: data.MerchantName,
	}, err
}
