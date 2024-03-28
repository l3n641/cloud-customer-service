package model

// MessageAddInput 更新商户信息
type MessageAddInput struct {
	Id            int
	TicketId      int    //工单id
	SenderType    int    // 发送类型
	ClientId      int    // 客户id
	ChatSupportId int    // 客户id
	Content       string // 内容
}
