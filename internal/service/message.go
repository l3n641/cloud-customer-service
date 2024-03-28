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
	IMessage interface {
		ClientSendMessage(ctx context.Context, in *model.MessageAddInput) (int64, error)
		GetMessageList(ctx context.Context, in *model.MessageListInput) (*model.ClientMessageListOutput, error)
		GetUnreadMessageQuantity(ctx context.Context, senderType, ticketId int) (int, error)
		ReadMessage(ctx context.Context, ticketId, senderType int) (int64, error)
	}
)

var (
	localMessage IMessage
)

func Message() IMessage {
	if localMessage == nil {
		panic("implement not found for interface IMessage, forgot register?")
	}
	return localMessage
}

func RegisterMessage(i IMessage) {
	localMessage = i
}
