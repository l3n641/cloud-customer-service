package message

import "github.com/gogf/gf/v2/frame/g"

type GetUnreadMessageQuantityReq struct {
	g.Meta `path:"/message-quantity" method:"get" summary:"获取未读取消息数" tags:"客户管理"`
}

type GetUnreadMessageQuantityRes struct {
	Quantity int `json:"quantity"`
}
