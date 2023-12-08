package admin

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sAdmin) GetUserInfo(ctx context.Context, id int) (user *entity.Admins, err error) {
	err = dao.Admins.Ctx(ctx).Where(g.Map{
		dao.Admins.Columns().Id: id,
	}).Scan(&user)
	return
}
