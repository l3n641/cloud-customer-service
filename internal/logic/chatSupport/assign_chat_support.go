package chatSupport

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"cloudCustomerService/internal/service"
	"context"
)

func (s *sChatSupport) AssignChatSupport(ctx context.Context, merchantId string) (*entity.ChatSupports, error) {
	var user *entity.ChatSupports

	merchant, err := service.Merchant().GetByMerchantId(ctx, merchantId)
	if merchant == nil || err != nil {
		return nil, err
	}

	where := dao.MerchantGroup.Ctx(ctx).As(dao.MerchantGroup.Table()).Builder().Where(dao.MerchantGroup.Table()+".merchant_id", merchant.Id)

	err = dao.ChatSupports.Ctx(ctx).LeftJoinOnFields(dao.MerchantGroup.Table(), dao.ChatSupports.Columns().Id, "=", dao.MerchantGroup.Columns().ChatSupportId).Where(
		where,
	).OrderDesc(dao.ChatSupports.Columns().IsOnline).Scan(&user)
	return user, err
}
