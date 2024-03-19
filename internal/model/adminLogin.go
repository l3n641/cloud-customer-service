package model

import "github.com/gogf/gf/v2/frame/g"

// AdminLoginInput 用户登录
type AdminLoginInput struct {
	Username string // 用户名
	Password string // 密码(明文)
	Ip       string //
}

type AdminLoginOutput struct {
	Data g.Map
}
