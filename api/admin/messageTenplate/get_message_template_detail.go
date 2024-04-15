package messageTenplate

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetMessageTemplateDetailReq struct {
	g.Meta `path:"/message-template/{id}" method:"get" summary:"查看消息模板详情" tags:"客服模块"`

	Id int `json:"id" dc:"消息模板id"`
}

type GetMessageTemplateDetailRes struct {
	Id      int    `json:"id"`
	MsgType int    `json:"msg_type" description:"消息类型"`
	Lang    string `json:"lang" description:"语种"`
	Content string `json:"content" description:"消息内容"`
}
