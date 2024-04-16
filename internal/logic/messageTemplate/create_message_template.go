package messageTemplate

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"cloudCustomerService/internal/service"
	"context"
)

type sMessageTemplate struct{}

func init() {
	service.RegisterMessageTemplate(New())
}

func New() *sMessageTemplate {
	return &sMessageTemplate{}
}

func (s *sMessageTemplate) CreateMessageTemplate(ctx context.Context, msgType int, lang, content string) (int64, error) {

	message := &entity.MessageTemplates{
		Content: content,
		MsgType: msgType,
		Lang:    lang,
	}
	result, err := dao.MessageTemplates.Ctx(ctx).Data().Insert(message)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()

}
