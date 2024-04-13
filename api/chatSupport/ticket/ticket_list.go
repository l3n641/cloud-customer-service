package ticket

import (
	"cloudCustomerService/api"
	"github.com/gogf/gf/v2/frame/g"
)

type SearchTicketReq struct {
	g.Meta `path:"/ticket" method:"get" summary:"查看工单" tags:"客服模块"`

	SearchFields
	api.PaginationReq
	api.OrderReq
}

type SearchFields struct {
	Account         string `json:"account" dc:"账号"`
	OnlyUnprocessed int    `json:"only_unprocessed" d:"0"  description:"只获取未处理的工单:0-获取全部,1-只查看未处理的"`
}

type SearchTicketRes struct {
	api.PaginationRes
}
