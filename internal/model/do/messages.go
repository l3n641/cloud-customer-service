// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Messages is the golang structure of table messages for DAO operations like Where/Data.
type Messages struct {
	g.Meta        `orm:"table:messages, do:true"`
	Id            interface{} //
	CreateAt      *gtime.Time // 消息时间
	TicketId      interface{} // 工单id
	SenderType    interface{} // 发送类型:1-客人发送给客服，2-客服发送给客人
	ClientId      interface{} // 顾客的id
	ChatSupportId interface{} // 客服的id
	Content       interface{} // 消息内容
	IsRead        interface{} // 消息是否被对方读取:0-未读，1-已读
	ReadTime      *gtime.Time // 读取消息的时间
}
