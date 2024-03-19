// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admins is the golang structure for table admins.
type Admins struct {
	Id          int         `json:"id"          ` //
	CreateAt    *gtime.Time `json:"createAt"    ` // 创建时间
	Username    string      `json:"username"    ` // 用户名
	Password    string      `json:"password"    ` // 加密后的密码
	Nickname    string      `json:"nickname"    ` // 昵称
	Avatar      string      `json:"avatar"      ` // 头像地址
	LastLoginAt *gtime.Time `json:"lastLoginAt" ` // 最后登录时间
	LastLoginIp string      `json:"lastLoginIp" ` // 最后登录ip
}
