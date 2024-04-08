package message

import "github.com/gogf/gf/v2/frame/g"

type AcknowledgeMessagesReq struct {
	g.Meta `path:"/msg-ack" method:"post" summary:"消息接收确认" tags:"客服管理"`

	TicketId   int   `v:"required" json:"ticket_id" dc:"工单id"`
	MessageIds []int `v:"required" json:"message_ids" dc:"消息id列表"`
}

type AcknowledgeMessagesRes struct {
}
