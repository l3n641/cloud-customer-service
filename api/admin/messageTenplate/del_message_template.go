package messageTenplate

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteMessageTemplateReq struct {
	g.Meta `path:"/message-template/{id}" method:"delete" summary:"删除消息模板" tags:"客服模块"`

	Id int `json:"id" dc:"消息模板id"`
}

type DeleteMessageTemplateRes struct {
}
