package chatSupport

import (
	"cloudCustomerService/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sChatSupport) UpdateChatSupportStatus(ctx context.Context, chatSupportId, status int) (row int64, err error) {

	data := g.Map{
		dao.ChatSupports.Columns().Status: status,
	}

	result, _ := dao.ChatSupports.Ctx(ctx).Data(data).Where(dao.ChatSupports.Columns().Id, chatSupportId).Update()
	return result.RowsAffected()

}
