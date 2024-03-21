package chatSupport

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
	"github.com/gogf/gf/v2/container/gmap"
)

func (s *sChatSupport) SearchChatSupport(ctx context.Context, in *model.ChatSupportSearchInput) (*model.ChatSupportSearchOutput, error) {

	out := &model.ChatSupportSearchOutput{
		Page: in.Page,
		Size: in.Size,
	}
	where := gmap.New()

	if in.Email != "" {
		where.Set(dao.ChatSupports.Columns().Email, in.Email)
	}

	if in.Status.String() != "" {
		where.Set(dao.ChatSupports.Columns().Status, in.Status.Int())
	}

	if in.IsOnline.String() != "" {
		where.Set(dao.ChatSupports.Columns().IsOnline, in.Status.Int())
	}

	query := dao.ChatSupports.Ctx(ctx).Where(where).Page(in.Page, in.Size)
	if in.OrderBy != "" {
		query = query.Order(in.OrderBy)
	}
	err := query.ScanAndCount(&out.List, &out.Total, false)
	return out, err

}
