package chatSupport

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
)

func (s *sChatSupport) GetDetailById(ctx context.Context, id int) (*model.ChatSupportDetail, error) {
	var user model.ChatSupportDetail
	err := dao.ChatSupports.Ctx(ctx).Where(dao.ChatSupports.Columns().Id, id).Scan(&user.ChatSupportItem)
	dao.MerchantGroup.Ctx(ctx).Where(dao.MerchantGroup.Columns().ChatSupportId, user.ChatSupportItem.Id).Scan(&user.MerchantGroup)
	return &user, err
}
