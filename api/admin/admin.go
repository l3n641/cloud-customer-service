// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package admin

import (
	"context"
	
	"cloudCustomerService/api/admin/account"
	"cloudCustomerService/api/admin/chatSupport"
	"cloudCustomerService/api/admin/merchant"
	"cloudCustomerService/api/admin/messageTenplate"
	"cloudCustomerService/api/admin/session"
)

type IAdminAccount interface {
	UserInfo(ctx context.Context, req *account.UserInfoReq) (res *account.UserInfoRes, err error)
}

type IAdminChatSupport interface {
	DetailChatSupport(ctx context.Context, req *chatSupport.DetailChatSupportReq) (res *chatSupport.DetailChatSupportRes, err error)
	SearchChatSupport(ctx context.Context, req *chatSupport.SearchChatSupportReq) (res *chatSupport.SearchChatSupportRes, err error)
	CreateChatSupport(ctx context.Context, req *chatSupport.CreateChatSupportReq) (res *chatSupport.CreateChatSupportRes, err error)
	UpdateChatSupport(ctx context.Context, req *chatSupport.UpdateChatSupportReq) (res *chatSupport.UpdateChatSupportRes, err error)
	UpdateChatSupportStatus(ctx context.Context, req *chatSupport.UpdateChatSupportStatusReq) (res *chatSupport.UpdateChatSupportStatusRes, err error)
}

type IAdminMerchant interface {
	CreateMerchant(ctx context.Context, req *merchant.CreateMerchantReq) (res *merchant.CreateMerchantRes, err error)
	DetailMerchant(ctx context.Context, req *merchant.DetailMerchantReq) (res *merchant.DetailMerchantRes, err error)
	SearchMerchant(ctx context.Context, req *merchant.SearchMerchantReq) (res *merchant.SearchMerchantRes, err error)
	UpdateMerchant(ctx context.Context, req *merchant.UpdateMerchantReq) (res *merchant.UpdateMerchantRes, err error)
}

type IAdminMessageTenplate interface {
	CreateMessageTemplate(ctx context.Context, req *messageTenplate.CreateMessageTemplateReq) (res *messageTenplate.CreateMessageTemplateRes, err error)
	DeleteMessageTemplate(ctx context.Context, req *messageTenplate.DeleteMessageTemplateReq) (res *messageTenplate.DeleteMessageTemplateRes, err error)
	GetMessageTemplateDetail(ctx context.Context, req *messageTenplate.GetMessageTemplateDetailReq) (res *messageTenplate.GetMessageTemplateDetailRes, err error)
	GetMessageTemplateList(ctx context.Context, req *messageTenplate.GetMessageTemplateListReq) (res *messageTenplate.GetMessageTemplateListRes, err error)
	UpdateMessageTemplate(ctx context.Context, req *messageTenplate.UpdateMessageTemplateReq) (res *messageTenplate.UpdateMessageTemplateRes, err error)
}

type IAdminSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
	SignOut(ctx context.Context, req *session.SignOutReq) (res *session.SignOutRes, err error)
}


