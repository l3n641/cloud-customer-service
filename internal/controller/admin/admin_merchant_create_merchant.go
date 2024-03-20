package admin

import (
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/merchant"
)

func (c *ControllerMerchant) CreateMerchant(ctx context.Context, req *merchant.CreateMerchantReq) (res *merchant.CreateMerchantRes, err error) {
	status, err := service.Merchant().CreateMerchant(ctx, &model.MerchantCreateInput{MerchantId: req.MerchantId, MerchantName: req.MerchantName})
	if status {
		return &merchant.CreateMerchantRes{}, nil
	}
	return nil, err
}
