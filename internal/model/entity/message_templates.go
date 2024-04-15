// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MessageTemplates is the golang structure for table message_templates.
type MessageTemplates struct {
	Id       int         `json:"id"       ` //
	CreateAt *gtime.Time `json:"createAt" ` // 消息时间
	MsgType  int         `json:"msgType"  ` // 消息类型 1-欢迎用语,2-离线用语,3-在线用语
	Lang     string      `json:"lang"     ` // 语种
	Content  string      `json:"content"  ` // 消息内容
}
