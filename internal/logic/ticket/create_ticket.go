package ticket

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

type sTicket struct{}

func init() {
	service.RegisterTicket(New())
}

func New() *sTicket {
	return &sTicket{}
}

func (s *sTicket) CreateTicket(ctx context.Context, clientId int, chatSupportId int) (int, error) {

	ticket := &entity.Tickets{
		ClientId:        clientId,
		ChatSupportId:   chatSupportId,
		LastMessageTime: gtime.Now(),
	}
	result, _ := dao.Tickets.Ctx(ctx).Data(ticket).Insert()
	id, err := result.LastInsertId()
	return int(id), err
}
