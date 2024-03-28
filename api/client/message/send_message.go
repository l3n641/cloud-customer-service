package message

import "github.com/gogf/gf/v2/frame/g"

type SendMessageReq struct {
	g.Meta `path:"/message" method:"post" summary:"客户发送消息" tags:"客户管理"`

	Content string `v:"required" json:"content" dc:"消息内容"`
}

type SendMessageRes struct {
}
