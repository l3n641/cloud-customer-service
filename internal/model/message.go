package model

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

// MessageAddInput 更新商户信息
type MessageAddInput struct {
	Id            int
	TicketId      int    //工单id
	SenderType    int    // 发送类型
	ClientId      int    // 客户id
	ChatSupportId int    // 客户id
	MsgType       int    // 消息类型
	Content       string // 内容
	RefererUrl    string // 客户当前所在页面
}

type ClientMessageListInput struct {
	Page          int // 分页号码
	Size          int // 分页数量，最大100
	TicketId      int // ticket id
	LastMessageId gvar.Var
	OnlyUnread    gvar.Var
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
	Id         int       `json:"id" description:"消息id"`
	CreateAt   time.Time `json:"create_at" description:"消息创建实际"`
	SenderType int       `json:"sender_type" description:"发送类型 1-顾客发送给客服,2-客服发送给客人"`
	MsgType    int       `json:"msg_type" description:"消息类型: 1-文本消息"`
	Content    string    `json:"content" description:"消息内容"`
	IsRead     int       `json:"is_read" description:"消息是否被对方读取:0-未读，1-已读"`
	ReadTime   time.Time `json:"read_time" description:"读取消息的时间"`
	RefererUrl string    `json:"referer_url" description:"客户发送消息时候所在的页面"`
}

type ChatSupportMessageItem struct {
	Id                           int         `json:"id" description:"消息id"`
	CreateAt                     *gtime.Time `json:"create_at" description:"消息创建实际"`
	TicketId                     int         `json:"ticket_id"      `  // 工单id
	ClientId                     int         `json:"client_id"      `  // 顾客的id
	ChatSupportId                int         `json:"chat_support_id" ` // 客服的id
	SenderType                   int         `json:"sender_type" description:"发送类型 1-顾客发送给客服,2-客服发送给客人"`
	MsgType                      int         `json:"msg_type" description:"消息类型: 1-文本消息"`
	Content                      string      `json:"content" description:"消息内容"`
	IsRead                       int         `json:"is_read" description:"消息是否被对方读取:0-未读，1-已读"`
	ReadTime                     *gtime.Time `json:"read_time" description:"读取消息的时间"`
	RefererUrl                   string      `json:"referer_url" description:"客户发送消息时候所在的页面"`
	ChatSupportMessageTicketItem `json:"ticket" description:"工单信息"`
}
type ChatSupportMessageTicketItem struct {
	Id               int         `json:"id"               `    //
	CreateAt         *gtime.Time `json:"create_at"         `   //
	ClientId         int         `json:"client_id"         `   // 客人id
	Account          string      `json:"account"          `    // 客户的账号名
	ChatSupportId    int         `json:"chat_support_id"    `  // 客服id
	LastMessageTime  *gtime.Time `json:"last_message_time"  `  // 最后发送消息时间
	CsUnreadMsgCount int         `json:"cs_unread_msg_count" ` // 客服未读消息数
	Remark           string      `json:"remark"           `    // 备注
}
