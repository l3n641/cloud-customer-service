// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admins is the golang structure of table admins for DAO operations like Where/Data.
type Admins struct {
	g.Meta      `orm:"table:admins, do:true"`
	Id          interface{} //
	CreateAt    *gtime.Time // 创建时间
	Username    interface{} // 用户名
	Password    interface{} // 加密后的密码
	Nickname    interface{} // 昵称
	Avatar      interface{} // 头像地址
	LastLoginAt *gtime.Time // 最后登录时间
	LastLoginIp interface{} // 最后登录ip
}
