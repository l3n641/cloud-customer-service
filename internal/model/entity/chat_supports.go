// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatSupports is the golang structure for table chat_supports.
type ChatSupports struct {
	Id          int         `json:"id"          ` //
	CreateAt    *gtime.Time `json:"createAt"    ` // 创建时间
	Email       string      `json:"email"       ` // 邮箱
	Password    string      `json:"password"    ` // 加密后的密码
	Nickname    string      `json:"nickname"    ` // 昵称
	Avatar      string      `json:"avatar"      ` // 头像地址
	LastLoginAt *gtime.Time `json:"lastLoginAt" ` // 最后登录时间
	LastLoginIp string      `json:"lastLoginIp" ` // 最后登录ip
	IsOnline    int         `json:"isOnline"    ` // 是否在线 1-在线 0-不在线
}
