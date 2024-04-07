package ticket

import (
	"cloudCustomerService/api"
	"github.com/gogf/gf/v2/frame/g"
)

type SearchTicketReq struct {
	g.Meta `path:"/ticket" method:"get" summary:"查看工单" tags:"客服管理"`

	SearchFields
	api.PaginationReq
	api.OrderReq
}

type SearchFields struct {
	Account string `json:"account" dc:"账号"`
}

type SearchTicketRes struct {
	api.PaginationRes
}
