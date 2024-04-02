package message

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"context"
)

func (s *sMessage) GetMessageById(ctx context.Context, id int64) (*entity.Messages, error) {

	var msg *entity.Messages
	err := dao.Messages.Ctx(ctx).Where(dao.Messages.Columns().Id, id).Scan(&msg)
	return msg, err

}
