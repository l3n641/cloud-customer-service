package chatSupport

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateChatSupportReq struct {
	g.Meta `path:"/chat-support" method:"post" summary:"创建客服" tags:"客服管理"`

	Email         string `v:"required#请输入邮箱" dc:"邮箱"`
	Nickname      string `dc:"昵称"`
	Password      string `v:"required#请输入密码" dc:"密码"`
	MerchantGroup []int  `v:"required#请选择要关联的商户" dc:"客服管理的商户"`
}

type CreateChatSupportRes struct {
}
