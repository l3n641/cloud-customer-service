package chatSupport

import (
	"cloudCustomerService/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sChatSupport) UpdateSessionInfo(ctx context.Context, id, isOnline int, ip string) error {

	data := g.Map{
		dao.ChatSupports.Columns().IsOnline:    isOnline,
		dao.ChatSupports.Columns().LastLoginAt: gtime.Now(),
		dao.ChatSupports.Columns().LastLoginIp: ip,
	}

	_, err := dao.ChatSupports.Ctx(ctx).Data(data).Where(dao.ChatSupports.Columns().Id, id).Update()
	return err

}
