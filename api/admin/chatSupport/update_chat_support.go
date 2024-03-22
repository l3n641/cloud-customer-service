package chatSupport

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateChatSupportReq struct {
	g.Meta `path:"/chat-support/{Id}" method:"put" summary:"更新客服" tags:"客服管理"`

	Id            int    `json:"id" dc:"客服id"`
	Nickname      string `dc:"昵称"`
	Password      string `dc:"密码"`
	MerchantGroup []int  `v:"required#请选择要关联的商户" dc:"客服管理的商户"`
}

type UpdateChatSupportRes struct {
}
