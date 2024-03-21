package admin

import (
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/merchant"
)

func (c *ControllerMerchant) UpdateMerchant(ctx context.Context, req *merchant.UpdateMerchantReq) (res *merchant.UpdateMerchantRes, err error) {
	status, err := service.Merchant().UpdateMerchant(ctx, &model.MerchantUpdateInput{
		Id:           req.Id,
		MerchantId:   req.MerchantId,
		MerchantName: req.MerchantName})
	if status {
		return &merchant.UpdateMerchantRes{}, nil
	}
	return nil, err

}
