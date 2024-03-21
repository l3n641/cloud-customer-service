package merchant

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sMerchant) UpdateMerchant(ctx context.Context, in *model.MerchantUpdateInput) (bool, error) {

	whereOr := dao.Merchants.Ctx(ctx).Builder().Where(g.Map{
		dao.Merchants.Columns().MerchantId: in.MerchantId,
	}).WhereOr(g.Map{
		dao.Merchants.Columns().MerchantName: in.MerchantName,
	})

	count, err := dao.Merchants.Ctx(ctx).Where(whereOr).Where("id <>", in.Id).CountColumn(dao.Merchants.Columns().Id)

	if err != nil {
		return false, err
	}

	if count > 0 {
		return false, gerror.New(`商户id或者商户名称已经存在`)
	}

	merchant := &entity.Merchants{
		Id:           in.Id,
		MerchantId:   in.MerchantId,
		MerchantName: in.MerchantName,
	}
	if _, err = dao.Merchants.Ctx(ctx).Data(merchant).Save(); err != nil {
		return false, err
	}

	return true, nil

}
