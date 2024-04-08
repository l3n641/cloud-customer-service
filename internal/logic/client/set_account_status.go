package client

import (
	"cloudCustomerService/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sClient) SetAccountStatus(ctx context.Context, clientId, status int) (int64, error) {
	result, _ := dao.Clients.Ctx(ctx).Where(g.Map{
		dao.Clients.Columns().Id: clientId,
	}).Data(g.Map{
		dao.Clients.Columns().Status:           status,
		dao.Clients.Columns().StatusChangeTime: gtime.Now(),
	}).Update()
	return result.RowsAffected()
}
