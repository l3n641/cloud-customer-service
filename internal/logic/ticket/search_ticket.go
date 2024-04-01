package ticket

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTicket) SearchTicket(ctx context.Context, in *model.TicketSearchInput) (*model.TicketSearchOutput, error) {

	out := &model.TicketSearchOutput{
		Page: in.Page,
		Size: in.Size,
	}
	where := g.Map{}
	if in.ChatSupportId != 0 {
		where[dao.Tickets.Columns().ChatSupportId] = in.ChatSupportId
	}

	order := fmt.Sprintf("%s desc,%s desc", dao.Tickets.Columns().CsUnreadMsgCount, dao.Tickets.Columns().LastMessageTime)
	err := dao.Tickets.Ctx(ctx).Where(where).Page(in.Page, in.Size).Order(order).ScanAndCount(&out.List, &out.Total, false)
	return out, err
}
