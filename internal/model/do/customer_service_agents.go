// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CustomerServiceAgents is the golang structure of table customer_service_agents for DAO operations like Where/Data.
type CustomerServiceAgents struct {
	g.Meta      `orm:"table:customer_service_agents, do:true"`
	Id          interface{} //
	CreateAt    *gtime.Time // 创建时间
	Email       interface{} // 邮箱
	Password    interface{} // 加密后的密码
	Nickname    interface{} // 昵称
	Avatar      interface{} // 头像地址
	LastLoginAt *gtime.Time // 最后登录时间
	LastLoginIp interface{} // 最后登录ip
	IsOnline    interface{} // 是否在线 1-在线 0-不在线
}
