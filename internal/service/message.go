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
	IMessage interface {
		GetMessageById(ctx context.Context, id int64) (*entity.Messages, error)
		GetMessageList(ctx context.Context, in *model.MessageListInput) (*model.ClientMessageListOutput, error)
		ClientGetMessageList(ctx context.Context, in *model.ClientMessageListInput) (*model.ClientMessageListOutput, error)
		GetUnreadMessageQuantity(ctx context.Context, senderType, ticketId int) (int, error)
		PushChatSupportMessage(ctx context.Context, messages *entity.Messages, ticket *entity.Tickets) (int64, error)
		ReadMessage(ctx context.Context, ticketId, senderType int) (int64, error)
		ClientReadAllMessage(ctx context.Context, ticketId, ChatSupportId int) (int64, error)
		SendMessage(ctx context.Context, in *model.MessageAddInput) (*entity.Messages, error)
		UpdateMessageReadStatus(ctx context.Context, ticketId int, senderType int, msgIds []int) (int64, error)
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
