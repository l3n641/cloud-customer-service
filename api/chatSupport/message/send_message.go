package message

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type SendMessageReq struct {
	g.Meta `path:"/message" method:"post" summary:"客户发送消息" tags:"客服模块"`

	TicketId int    `v:"required" json:"ticket_id" dc:"工单id"`
	Content  string `v:"required" json:"content" dc:"消息内容"`
	MsgType  int    `v:"required|in:1" json:"msg_type" dc:"消息类型 1-文本消息"`
}

type SendMessageRes struct {
	Id         int         `json:"id" description:"消息id"`
	CreateAt   *gtime.Time `json:"create_at" description:"消息创建实际"`
	TicketId   int         `json:"ticket_id" description:"工单id"`
	ClientId   int         `json:"client_id" description:"顾客id"`
	SenderType int         `json:"sender_type" description:"发送类型 1-顾客发送给客服,2-客服发送给客人"`
	MsgType    int         `json:"msg_type" description:"消息类型: 1-文本消息"`
	Content    string      `json:"content" description:"消息内容"`
	IsRead     int         `json:"is_read" description:"消息是否被对方读取:0-未读，1-已读"`
	ReadTime   *gtime.Time `json:"read_time" description:"读取消息的时间"`
}
