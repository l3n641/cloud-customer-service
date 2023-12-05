package session

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type LoginReq struct {
	g.Meta `path:"/login" method:"post" summary:"管理员登录" tags:"管理员登录"`

	Email    string `v:"required#请输入邮箱" dc:"邮箱"`
	Password string `v:"required#请输入密码" dc:"密码"`
}

type LoginRes struct {
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}
