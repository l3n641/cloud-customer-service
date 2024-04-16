package model

import "github.com/gogf/gf/v2/container/gvar"

type MessageTemplateListInput struct {
	Pagination
	MsgType gvar.Var
	Lang    string
}

type MessageTemplateListOutput struct {
	Total int                   `json:"total" description:"数据总数"`
	List  []MessageTemplateItem `json:"list" description:"列表"`
}

type MessageTemplateItem struct {
	Id      int    `json:"id"`
	MsgType int    `json:"msg_type" description:"消息类型"`
	Lang    string `json:"lang" description:"语种"`
	Content string `json:"content" description:"消息内容"`
}
