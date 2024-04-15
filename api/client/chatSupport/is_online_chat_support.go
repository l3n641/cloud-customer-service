package chatSupport

import (
	"cloudCustomerService/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type IsOnlineChatSupportReq struct {
	g.Meta `path:"/chat-support" method:"get" summary:"查看客服是否在线" tags:"客户管理"`
}

type IsOnlineChatSupportRes struct {
	IsOnline int                      `json:"is_online" description:"客服是否在线 -1-没有对应的客服 0-不在线 1-在线 "`
	Messages []model.MessageTemplates `json:"messages" description:"消息模板列表"`
}
