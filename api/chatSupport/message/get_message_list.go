package message

import (
	"cloudCustomerService/api"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

type GetMessageListReq struct {
	g.Meta `path:"/message/" method:"get" summary:"查看消息列表" tags:"客服模块"`

	TicketId      int      `v:"required" json:"ticket_id" dc:"工单id"`
	LastMessageId gvar.Var `json:"last_message_id" description:"最后一条消息id"`
	api.PaginationReq
}

type GetMessageListRes struct {
	api.PaginationRes
}
