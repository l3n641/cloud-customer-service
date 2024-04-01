package account

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserInfoReq struct {
	g.Meta `path:"/user-info" method:"get" summary:"获取客服账号信息" tags:"客服模块"`
}

type UserInfoRes struct {
	Email    string `json:"email"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}
