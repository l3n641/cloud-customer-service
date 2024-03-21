package account

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserInfoReq struct {
	g.Meta `path:"/admin-info" method:"get" summary:"获取管理员账号信息" tags:"管理员简介管理"`
}

type UserInfoRes struct {
	Username string `json:"username"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}
