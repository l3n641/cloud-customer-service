package client

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	sClient struct{}
)

func init() {
	service.RegisterClient(New())
}
func New() *sClient {
	return &sClient{}
}

func (s *sClient) GetIdentityKey() string {
	return "client_id"
}

func (s *sClient) createClient(ctx context.Context, email, phone, domain, merchantId string) (userId int64, err error) {
	user := &entity.Clients{
		Email:      email,
		Phone:      phone,
		Domain:     domain,
		MerchantId: merchantId,
	}

	result, err := dao.Clients.Ctx(ctx).Data(user).Insert()
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return id, err
}

func (s *sClient) Login(ctx context.Context, in *model.ClientLoginInput) (*model.ClientLoginOutput, error) {
	where := g.Map{
		dao.Clients.Columns().Domain: in.Domain,
	}

	if in.LoginType == consts.LoginTypeEmail {
		where[dao.Clients.Columns().Email] = in.Email
	} else {
		where[dao.Clients.Columns().Phone] = in.Phone
	}
	var client *entity.Clients

	err := dao.Clients.Ctx(ctx).Where(where).Scan(&client)

	if err != nil {
		return nil, err
	}
	var userId int64
	if client == nil {
		if userId, err = s.createClient(ctx, in.Email, in.Phone, in.Domain, in.MerchantId); err != nil {
			return nil, err
		}

	} else {
		userId = int64(client.Id)
	}

	dao.Clients.Ctx(ctx).Data(g.Map{
		dao.Clients.Columns().Ip:            in.Ip,
		dao.Clients.Columns().LastLoginTime: gtime.Now(),
		dao.Clients.Columns().UserAgent:     in.UserAgent,
		dao.Clients.Columns().Lang:          in.Lang,
	}).Where(g.Map{
		dao.Clients.Columns().Id: userId,
	}).Update()

	return &model.ClientLoginOutput{
		Data: g.Map{
			s.GetIdentityKey(): userId,
		},
	}, nil
}
