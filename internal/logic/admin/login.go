package admin

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type (
	sAdmin struct{}
)

func init() {
	admin := NewAdmin()
	service.RegisterAdmin(admin)
}
func NewAdmin() *sAdmin {
	return &sAdmin{}
}

func (s *sAdmin) encryptPassword(password string) string {
	return gmd5.MustEncrypt(password)
}

func (s *sAdmin) GetIdentityKey() string {
	return "admin_id"
}

func (s *sAdmin) getUserByPassportAndPassword(ctx context.Context, username, password string) (user *entity.Admins, err error) {
	err = dao.Admins.Ctx(ctx).Where(g.Map{
		dao.Admins.Columns().Username: username,
		dao.Admins.Columns().Password: s.encryptPassword(password),
	}).Scan(&user)
	return
}

func (s *sAdmin) Login(ctx context.Context, in *model.AdminLoginInput) (*model.AdminLoginOutput, error) {

	userEntity, err := s.getUserByPassportAndPassword(
		ctx,
		in.Username,
		in.Password,
	)

	if err != nil {
		return nil, err
	}

	if userEntity == nil {
		return nil, gerror.New(`账号或密码错误`)
	}

	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")

	dao.Admins.Ctx(ctx).Data(g.Map{
		dao.Admins.Columns().LastLoginIp: in.Ip,
		dao.Admins.Columns().LastLoginAt: timeString,
	}).Where(g.Map{
		dao.Admins.Columns().Id: userEntity.Id,
	}).Update()

	return &model.AdminLoginOutput{
		Data: g.Map{
			s.GetIdentityKey(): userEntity.Id,
		},
	}, nil
}
