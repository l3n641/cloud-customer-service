package session

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type LoginReq struct {
	g.Meta `path:"/session" method:"post" summary:"客服账户登录" tags:"客服模块"`

	Email    string `v:"required|email" json:"email" dc:"邮箱"`
	Password string `v:"required" json:"password" dc:"密码"`
}

type LoginRes struct {
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}
