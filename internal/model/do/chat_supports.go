// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatSupports is the golang structure of table chat_supports for DAO operations like Where/Data.
type ChatSupports struct {
	g.Meta      `orm:"table:chat_supports, do:true"`
	Id          interface{} //
	CreateAt    *gtime.Time // 创建时间
	Email       interface{} // 邮箱
	Password    interface{} // 加密后的密码
	Nickname    interface{} // 昵称
	Avatar      interface{} // 头像地址
	LastLoginAt *gtime.Time // 最后登录时间
	LastLoginIp interface{} // 最后登录ip
	IsOnline    interface{} // 是否在线 1-在线 0-不在线
	Status      interface{} // 账户状态 0-冻结 1-正常
}
