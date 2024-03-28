package client

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sClient) GetClientById(ctx context.Context, id int) (*entity.Clients, error) {

	var client *entity.Clients

	err := dao.Clients.Ctx(ctx).Where(g.Map{
		dao.Clients.Columns().Id: id,
	}).Scan(&client)
	return client, err
}
