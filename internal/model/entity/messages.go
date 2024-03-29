// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Messages is the golang structure for table messages.
type Messages struct {
	Id            int         `json:"id"            ` //
	CreateAt      *gtime.Time `json:"createAt"      ` // 消息时间
	TicketId      int         `json:"ticketId"      ` // 工单id
	SenderType    int         `json:"senderType"    ` // 发送类型:1-客人发送给客服，2-客服发送给客人
	ClientId      int         `json:"clientId"      ` // 顾客的id
	ChatSupportId int         `json:"chatSupportId" ` // 客服的id
	MsgType       int         `json:"msgType"       ` // 消息类型 1-文本消息
	Content       string      `json:"content"       ` // 消息内容
	IsRead        int         `json:"isRead"        ` // 消息是否被对方读取:0-未读，1-已读
	ReadTime      *gtime.Time `json:"readTime"      ` // 读取消息的时间
	RefererUrl    string      `json:"refererUrl"    ` // 客户发送消息时候所在的页面
}
