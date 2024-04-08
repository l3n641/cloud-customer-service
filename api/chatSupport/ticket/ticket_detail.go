package ticket

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type DetailTicketReq struct {
	g.Meta `path:"/ticket/{Id}" method:"get" summary:"获取工单详细信息" tags:"客服模块"`
	Id     int `json:"id" dc:"工单id"`
}

type DetailTicketRes struct {
	DetailTicketInfo `json:"detail_ticket_info"`
	ClientInfo       `json:"client_info"`
}

type DetailTicketInfo struct {
	Id              int         `json:"id"`
	CreateAt        *gtime.Time `json:"create_at"`
	LastMessageTime *gtime.Time `json:"last_message_time"`
	Remark          string      `json:"remark"`
}

type ClientInfo struct {
	Id                 int    `json:"id"` // id
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	Domain             string `json:"domain"`
	MerchantId         string `json:"merchant_id"`
	Ip                 string `json:"ip"`
	BrowserFingerprint string `json:"browser_fingerprint"`
	Area               string `json:"area"`
	UserAgent          string `json:"user_agent"`
	Lang               string `json:"lang"`
}
