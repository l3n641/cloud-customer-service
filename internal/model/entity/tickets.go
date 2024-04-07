// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tickets is the golang structure for table tickets.
type Tickets struct {
	Id               int         `json:"id"               ` //
	CreateAt         *gtime.Time `json:"createAt"         ` //
	ClientId         int         `json:"clientId"         ` // 客人id
	Account          string      `json:"account"          ` // 客户的账号名
	ChatSupportId    int         `json:"chatSupportId"    ` // 客服id
	LastMessageTime  *gtime.Time `json:"lastMessageTime"  ` // 最后发送消息时间
	CsUnreadMsgCount int         `json:"csUnreadMsgCount" ` // 客服未读消息数
	Remark           string      `json:"remark"           ` // 备注
}
