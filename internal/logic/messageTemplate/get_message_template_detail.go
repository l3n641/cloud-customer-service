package messageTemplate

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sMessageTemplate) GetMessageDetail(ctx context.Context, id int) (*entity.MessageTemplates, error) {
	var data *entity.MessageTemplates
	err := dao.MessageTemplates.Ctx(ctx).Where(g.Map{
		dao.MessageTemplates.Columns().Id: id,
	}).Scan(&data)

	return data, err

}
