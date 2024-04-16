package messageTenplate

import "github.com/gogf/gf/v2/frame/g"

type CreateMessageTemplateReq struct {
	g.Meta `path:"/message-template" method:"post" summary:"创建消息模板" tags:"客服模块"`

	MsgType int    `v:"required|in:1,2,3" json:"msg_type" dc:"消息类型 1-欢迎用语,2-离线用语,3-在线用语"`
	Content string `v:"required" json:"content" dc:"消息内容"`
	Lang    string `v:"required|length:2,5" json:"lang" dc:"语种"`
}

type CreateMessageTemplateRes struct {
	Id int64 `json:"id" description:"id"`
}
