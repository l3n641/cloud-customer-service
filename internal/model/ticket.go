package model

import (
	"cloudCustomerService/internal/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
)

// TicketSearchInput 搜索工单input
type TicketSearchInput struct {
	Page          int // 分页号码
	Size          int // 分页数量，最大100
	ChatSupportId int //客服id
	SearchTicketFields
}

type SearchTicketFields struct {
}

// TicketSearchOutput 查询列表结果
type TicketSearchOutput struct {
	List  []TicketItem `json:"list" description:"列表"`
	Page  int          `json:"page" description:"分页码"`
	Size  int          `json:"size" description:"分页数量"`
	Total int          `json:"total" description:"数据总数"`
}

type TicketItem struct {
	Id               int         `json:"id"`
	CreateAt         *gtime.Time `json:"create_at"`
	ClientId         int         `json:"client_id"`
	ChatSupportId    int         `json:"chat_support_id"`
	LastMessageTime  *gtime.Time `json:"last_message_time"`
	CsUnreadMsgCount int         `json:"cs_unread_msg_count"`
}

type TicketDetail struct {
	Ticket *entity.Tickets
	Client *entity.Clients
}
