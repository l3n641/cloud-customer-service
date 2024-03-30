package chatSupport

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sChatSupport) getUserByPassportAndPassword(ctx context.Context, email, password string) (user *entity.Admins, err error) {
	err = dao.ChatSupports.Ctx(ctx).Where(g.Map{
		dao.ChatSupports.Columns().Email:    email,
		dao.ChatSupports.Columns().Password: s.encryptPassword(password),
	}).Scan(&user)
	return
}

func (s *sChatSupport) Login(ctx context.Context, in *model.ChatSupportLoginInput) (*model.ChatSupportLoginOutput, error) {

	userEntity, err := s.getUserByPassportAndPassword(
		ctx,
		in.Email,
		in.Password,
	)

	if err != nil {
		return nil, err
	}

	if userEntity == nil {
		return nil, gerror.New(`账号或密码错误`)
	}

	dao.ChatSupports.Ctx(ctx).Data(g.Map{
		dao.ChatSupports.Columns().LastLoginIp: in.Ip,
		dao.ChatSupports.Columns().LastLoginAt: gtime.Now(),
	}).Where(g.Map{
		dao.Admins.Columns().Id: userEntity.Id,
	}).Update()

	return &model.ChatSupportLoginOutput{
		Data: g.Map{
			s.GetIdentityKey(): userEntity.Id,
		},
	}, nil
}
