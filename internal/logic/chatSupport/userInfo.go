package chatSupport

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sChatSupport) GetUserInfo(ctx context.Context, id int) (user *entity.ChatSupports, err error) {
	err = dao.ChatSupports.Ctx(ctx).Where(g.Map{
		dao.ChatSupports.Columns().Id: id,
	}).Scan(&user)
	return
}
