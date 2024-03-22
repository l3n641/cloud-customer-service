package chatSupport

import (
	"cloudCustomerService/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type DetailChatSupportReq struct {
	g.Meta `path:"/chat-support/{Id}" method:"get" summary:"获取客服详细信息" tags:"客服管理"`
	Id     int `json:"id" dc:"客服id"`
}

type DetailChatSupportRes struct {
	Id            int                          `json:"id"        `                     // id
	Email         string                       `json:"email"            `              // 客服账户
	MerchantGroup []*model.ChatSupportMerchant `json:"merchant_group"                ` //客服所属商户
}
