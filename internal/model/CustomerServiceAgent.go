package model

import "github.com/gogf/gf/v2/frame/g"

// CreateCustomerServiceAgentInput 创建客服
type CreateCustomerServiceAgentInput struct {
	Email    string // 账号
	Password string // 密码(明文)
	NickName string // 昵称
}

// CreateCustomerServiceAgentOutput 创建客服
type CreateCustomerServiceAgentOutput struct {
	Data g.Map
}
