// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"context"
)

type (
	IChatSupport interface {
		AssignChatSupport(ctx context.Context, merchantId string) (*entity.ChatSupports, error)
		GetIdentityKey() string
		// CreteChatSupport 创建客服
		CreteChatSupport(ctx context.Context, in *model.ChatSupportCreateInput) (err error)
		GetDetailById(ctx context.Context, id int) (*model.ChatSupportDetail, error)
		GetOnlineStatus(ctx context.Context, id int) (*model.IsOnlineChatSupport, error)
		Login(ctx context.Context, in *model.ChatSupportLoginInput) (*model.ChatSupportLoginOutput, error)
		SearchChatSupport(ctx context.Context, in *model.ChatSupportSearchInput) (*model.ChatSupportSearchOutput, error)
		UpdateChatSupport(ctx context.Context, in *model.ChatSupportUpdateInput) error
		UpdateSessionInfo(ctx context.Context, id, isOnline int, ip string) error
		GetUserInfo(ctx context.Context, id int) (user *entity.ChatSupports, err error)
	}
)

var (
	localChatSupport IChatSupport
)

func ChatSupport() IChatSupport {
	if localChatSupport == nil {
		panic("implement not found for interface IChatSupport, forgot register?")
	}
	return localChatSupport
}

func RegisterChatSupport(i IChatSupport) {
	localChatSupport = i
}
