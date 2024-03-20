package merchant

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sMerchant struct{}

func init() {
	merchant := NewMerchant()
	service.RegisterMerchant(merchant)
}

func NewMerchant() *sMerchant {
	return &sMerchant{}
}

func (s *sMerchant) CreateMerchant(ctx context.Context, in *model.MerchantCreateInput) (bool, error) {
	count, err := dao.Merchants.Ctx(ctx).Where(g.Map{
		dao.Merchants.Columns().MerchantId: in.MerchantId,
	}).WhereOr(g.Map{
		dao.Merchants.Columns().MerchantName: in.MerchantName,
	}).CountColumn(dao.Merchants.Columns().Id)

	if err != nil {
		return false, err
	}

	if count > 0 {
		return false, gerror.New(`商户id或者商户名称已经存在`)
	}

	merchant := &entity.Merchants{
		MerchantId:   in.MerchantId,
		MerchantName: in.MerchantName,
	}
	_, err = dao.Merchants.Ctx(ctx).Data(merchant).Insert()
	return true, nil
}
