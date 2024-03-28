package merchant

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"context"
)

func (s *sMerchant) GetByMerchantId(ctx context.Context, merchantId string) (*entity.Merchants, error) {

	m := entity.Merchants{}
	err := dao.Merchants.Ctx(ctx).Where(dao.Merchants.Columns().MerchantId, merchantId).Scan(&m)
	return &m, err
}
