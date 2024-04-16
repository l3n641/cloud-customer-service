package chatSupport

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"cloudCustomerService/internal/service"
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
)

type sChatSupport struct{}

func init() {
	service.RegisterChatSupport(New())
}

func New() *sChatSupport {
	return &sChatSupport{}
}

func (s *sChatSupport) GetIdentityKey() string {
	return "chat_support_id"
}

func (s *sChatSupport) isEmailExist(ctx context.Context, email string) (bool, error) {
	count, err := dao.ChatSupports.Ctx(ctx).Where(g.Map{
		dao.ChatSupports.Columns().Email: email,
	}).Count()
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (s *sChatSupport) encryptPassword(password string) string {
	return gmd5.MustEncrypt(password)
}

// CreteChatSupport 创建客服
func (s *sChatSupport) CreteChatSupport(ctx context.Context, in *model.ChatSupportCreateInput) (err error) {
	isExist, err := s.isEmailExist(ctx, in.Email)
	if err != nil {
		return err
	}

	if isExist {
		return errors.New("邮箱已经存在")
	}

	user := &entity.ChatSupports{
		Email:    in.Email,
		Password: s.encryptPassword(in.Password),
		Nickname: in.NickName,
		Status:   consts.ChatSupportActive,
	}

	result, err := dao.ChatSupports.Ctx(ctx).Data(user).Insert()
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	groups := g.List{}
	if err != nil {
		return err
	}

	for _, merchant := range in.MerchantGroup {
		groups = append(groups, g.Map{
			dao.MerchantGroup.Columns().ChatSupportId: userId,
			dao.MerchantGroup.Columns().MerchantId:    merchant,
		})
	}

	_, err = dao.MerchantGroup.Ctx(ctx).Data(groups).Insert()
	if err != nil {
		return err
	}

	return err
}
