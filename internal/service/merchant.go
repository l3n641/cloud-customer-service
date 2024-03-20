// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"cloudCustomerService/internal/model"
	"context"
)

type (
	IMerchant interface {
		CreateMerchant(ctx context.Context, in *model.MerchantCreateInput) (bool, error)
	}
)

var (
	localMerchant IMerchant
)

func Merchant() IMerchant {
	if localMerchant == nil {
		panic("implement not found for interface IMerchant, forgot register?")
	}
	return localMerchant
}

func RegisterMerchant(i IMerchant) {
	localMerchant = i
}
