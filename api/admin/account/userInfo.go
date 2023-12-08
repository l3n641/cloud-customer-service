package account

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserInfoReq struct {
	g.Meta `path:"/user-info" method:"get" summary:"管理员账号信息" tags:"管理员"`
}

type UserInfoRes struct {
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}
