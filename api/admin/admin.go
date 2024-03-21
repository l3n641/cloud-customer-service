// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package admin

import (
	"context"
	
	"cloudCustomerService/api/admin/account"
	"cloudCustomerService/api/admin/merchant"
	"cloudCustomerService/api/admin/session"
)

type IAdminAccount interface {
	UserInfo(ctx context.Context, req *account.UserInfoReq) (res *account.UserInfoRes, err error)
}

type IAdminMerchant interface {
	CreateMerchant(ctx context.Context, req *merchant.CreateMerchantReq) (res *merchant.CreateMerchantRes, err error)
	DetailMerchant(ctx context.Context, req *merchant.DetailMerchantReq) (res *merchant.DetailMerchantRes, err error)
	SearchMerchant(ctx context.Context, req *merchant.SearchMerchantReq) (res *merchant.SearchMerchantRes, err error)
	UpdateMerchant(ctx context.Context, req *merchant.UpdateMerchantReq) (res *merchant.UpdateMerchantRes, err error)
}

type IAdminSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
}


