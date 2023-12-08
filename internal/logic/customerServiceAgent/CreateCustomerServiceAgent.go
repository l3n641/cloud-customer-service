package customerServiceAgent

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"cloudCustomerService/internal/service"
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
)

type sCustomerServiceAgent struct{}

func init() {
	service.RegisterCustomerServiceAgent(New())
}

func New() *sCustomerServiceAgent {
	return &sCustomerServiceAgent{}
}

func (s *sCustomerServiceAgent) isEmailExist(ctx context.Context, email string) (bool, error) {
	count, err := dao.CustomerServiceAgents.Ctx(ctx).Where(g.Map{
		dao.Admins.Columns().Email: email,
	}).Count()
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (s *sCustomerServiceAgent) encryptPassword(password string) string {
	return gmd5.MustEncrypt(password)
}

func (s *sCustomerServiceAgent) Crete(ctx context.Context, in *model.CreateCustomerServiceAgentInput) (err error) {
	isExist, err := s.isEmailExist(ctx, in.Email)
	if err != nil {
		return err
	}

	if isExist {
		return errors.New("邮箱已经存在")
	}

	user := &entity.CustomerServiceAgents{
		Email:    in.Email,
		Password: s.encryptPassword(in.Password),
		Nickname: in.NickName,
	}
	_, err = dao.CustomerServiceAgents.Ctx(ctx).Data(user).Insert()
	return err
}
