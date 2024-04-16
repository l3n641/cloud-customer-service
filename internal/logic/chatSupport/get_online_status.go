package chatSupport

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
)

func (s *sChatSupport) GetOnlineStatus(ctx context.Context, id int) (*model.IsOnlineChatSupport, error) {

	data := &model.IsOnlineChatSupport{
		Messages: make([]model.MessageTemplates, 0),
	}
	err := dao.ChatSupports.Ctx(ctx).Fields(dao.ChatSupports.Columns().IsOnline).Where(dao.ChatSupports.Columns().Id, id).Scan(&data)
	if err != nil {
		return nil, err
	}
	query := dao.MessageTemplates.Ctx(ctx)
	if data.IsOnline == 1 {
		//在线消息模板
		query = query.Where(dao.MessageTemplates.Columns().MsgType, 3)
	} else {
		//离线消息模板
		query = query.Where(dao.MessageTemplates.Columns().MsgType, 2)
	}
	err = query.Scan(&data.Messages)

	return data, err
}
