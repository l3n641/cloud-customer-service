package message

import "github.com/gogf/gf/v2/frame/g"

type SendMessageReq struct {
	g.Meta `path:"/message" method:"post" summary:"客户发送消息" tags:"客户管理"`

	Content    string `v:"required" json:"content" dc:"消息内容"`
	MsgType    int    `v:"required|in:1" json:"msg_type" dc:"消息类型 1-文本消息"`
	RefererUrl string `json:"referer_url" dc:"客户发送消息时候所在的页面"`
}

type SendMessageRes struct {
}
