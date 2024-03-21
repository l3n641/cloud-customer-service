package merchant

import (
	"cloudCustomerService/api"
	"github.com/gogf/gf/v2/frame/g"
)

type SearchMerchantReq struct {
	g.Meta `path:"/merchant" method:"get" summary:"查看商户" tags:"管理员管理商户"`

	SearchFields
	api.PaginationReq
}

type SearchFields struct {
}

type SearchMerchantRes struct {
	api.PaginationRes
}
