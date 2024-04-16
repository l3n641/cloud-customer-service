// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MessageTemplates is the golang structure of table message_templates for DAO operations like Where/Data.
type MessageTemplates struct {
	g.Meta   `orm:"table:message_templates, do:true"`
	Id       interface{} //
	CreateAt *gtime.Time // 消息时间
	MsgType  interface{} // 消息类型 1-欢迎用语,2-离线用语,3-在线用语
	Lang     interface{} // 语种
	Content  interface{} // 消息内容
}
