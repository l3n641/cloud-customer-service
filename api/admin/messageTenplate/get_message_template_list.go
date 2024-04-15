package messageTenplate

import (
	"cloudCustomerService/api"
	"github.com/gogf/gf/v2/frame/g"
)

type GetMessageTemplateListReq struct {
	g.Meta `path:"/message-template" method:"get" summary:"查看消息模板列表" tags:"客服模块"`

	MsgType g.Var  ` json:"msg_type" dc:"消息类型 1-欢迎用语,2-离线用语,3-在线用语"`
	Lang    string ` json:"lang" dc:"语种"`
	api.PaginationReq
}

type GetMessageTemplateListRes struct {
	api.PaginationRes
}
