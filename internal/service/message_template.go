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
	IMessageTemplate interface {
		CreateMessageTemplate(ctx context.Context, msgType int, lang, content string) (int64, error)
		DeleteMessage(ctx context.Context, id int) (int64, error)
		GetMessageDetail(ctx context.Context, id int) (*entity.MessageTemplates, error)
		GetMessageList(ctx context.Context, in *model.MessageTemplateListInput) (*model.MessageTemplateListOutput, error)
		GetWelcomeMessage(ctx context.Context) (*[]model.MessageTemplateItem, error)
		UpdateMessageTemplate(ctx context.Context, id, msgType int, lang, content string) (int64, error)
	}
)

var (
	localMessageTemplate IMessageTemplate
)

func MessageTemplate() IMessageTemplate {
	if localMessageTemplate == nil {
		panic("implement not found for interface IMessageTemplate, forgot register?")
	}
	return localMessageTemplate
}

func RegisterMessageTemplate(i IMessageTemplate) {
	localMessageTemplate = i
}
