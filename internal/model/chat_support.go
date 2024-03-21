package model

// ChatSupportCreateInput 创建客服输入
type ChatSupportCreateInput struct {
	Email         string // 账号
	Password      string // 密码(明文)
	NickName      string // 昵称
	MerchantGroup []int  //要关联的商户

}

// ChatSupportCreateOutput 创建客服输出
type ChatSupportCreateOutput struct {
}
