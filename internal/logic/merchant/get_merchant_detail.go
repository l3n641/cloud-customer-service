package merchant

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"context"
)

func (s *sMerchant) GetDetailById(ctx context.Context, id int) (*entity.Merchants, error) {

	m := entity.Merchants{}
	err := dao.Merchants.Ctx(ctx).Where(dao.Merchants.Columns().Id, id).Scan(&m)
	return &m, err
}
