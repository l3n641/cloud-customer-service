package customerServiceAgent

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateCustomerServiceAgentsReq struct {
	g.Meta `path:"/customer-service-agent" method:"post" summary:"创建客服" tags:"客服"`

	Email    string `v:"required#请输入邮箱" dc:"邮箱"`
	Nickname string `v:"required#请输入昵称" dc:"昵称"`
	Password string `v:"required#请输入密码" dc:"密码"`
}

type CreateCustomerServiceAgentsRes struct {
}
