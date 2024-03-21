package chatSupport

import (
	"cloudCustomerService/api"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

type SearchChatSupportReq struct {
	g.Meta `path:"/chat-support" method:"get" summary:"查看商户" tags:"客服管理"`

	SearchFields
	api.PaginationReq
	api.OrderReq
}

type SearchFields struct {
	Email    string   `json:"email"`     //客服邮箱
	Status   gvar.Var `json:"status"`    //客服状态
	IsOnline gvar.Var `json:"is_online"` //客服是否在线
}

type SearchChatSupportRes struct {
	api.PaginationRes
}
