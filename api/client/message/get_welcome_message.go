package message

import (
	"cloudCustomerService/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetWelcomeMessageReq struct {
	g.Meta `path:"/welcome-message" method:"get" summary:"获取欢迎用语" tags:"客户管理"`
}

type GetWelcomeMessageRes struct {
	Item *[]model.MessageTemplateItem
}
