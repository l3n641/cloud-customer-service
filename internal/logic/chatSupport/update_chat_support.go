package chatSupport

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sChatSupport) UpdateChatSupport(ctx context.Context, in *model.ChatSupportUpdateInput) error {

	data := g.Map{
		dao.ChatSupports.Columns().Nickname: in.NickName,
	}

	if in.Password != "" {
		data[dao.ChatSupports.Columns().Password] = s.encryptPassword(in.Password)

	}

	if _, err := dao.ChatSupports.Ctx(ctx).Data(data).Where(dao.ChatSupports.Columns().Id, in.Id).Update(); err != nil {
		return err
	}

	groups := g.List{}
	for _, merchant := range in.MerchantGroup {
		groups = append(groups, g.Map{
			dao.MerchantGroup.Columns().ChatSupportId: in.Id,
			dao.MerchantGroup.Columns().MerchantId:    merchant,
		})
	}

	dao.MerchantGroup.Ctx(ctx).Where(dao.MerchantGroup.Columns().ChatSupportId, in.Id).Delete()
	if _, err := dao.MerchantGroup.Ctx(ctx).Data(groups).Insert(); err != nil {
		return err
	}

	return nil

}
