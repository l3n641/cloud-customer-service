// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tickets is the golang structure of table tickets for DAO operations like Where/Data.
type Tickets struct {
	g.Meta           `orm:"table:tickets, do:true"`
	Id               interface{} //
	CreateAt         *gtime.Time //
	ClientId         interface{} // 客人id
	Account          interface{} // 客户的账号名
	ChatSupportId    interface{} // 客服id
	LastMessageTime  *gtime.Time // 最后发送消息时间
	CsUnreadMsgCount interface{} // 客服未读消息数
	Remark           interface{} // 备注
}
