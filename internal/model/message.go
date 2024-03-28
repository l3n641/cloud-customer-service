package model

import "github.com/gogf/gf/v2/container/gvar"

// MessageAddInput 更新商户信息
type MessageAddInput struct {
	Id            int
	TicketId      int    //工单id
	SenderType    int    // 发送类型
	ClientId      int    // 客户id
	ChatSupportId int    // 客户id
	Content       string // 内容
}

// MessageListInput 消息列表
type MessageListInput struct {
	Page          int // 分页号码
	Size          int // 分页数量，最大100
	TicketId      int // ticket id
	LastMessageId gvar.Var
}

// ClientMessageListOutput 查询列表结果
type ClientMessageListOutput struct {
	List  []ClientMessageItem `json:"list" description:"列表"`
	Page  int                 `json:"page" description:"分页码"`
	Size  int                 `json:"size" description:"分页数量"`
	Total int                 `json:"total" description:"数据总数"`
}

type ClientMessageItem struct {
	Id         int    `json:"id" description:"消息id"`
	SenderType int    `json:"sender_type" description:"发送类型 1-顾客发送给客服,2-客服发送给客人"`
	Content    string `json:"content" description:"消息内容"`
}
