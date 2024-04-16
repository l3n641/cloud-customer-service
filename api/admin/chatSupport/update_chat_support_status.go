package chatSupport

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateChatSupportStatusReq struct {
	g.Meta `path:"/chat-support-status/{Id}" method:"put" summary:"更新客服状态" tags:"管理员模块"`

	Id     int `json:"id" dc:"客服id"`
	Status int `json:"status" v:"required|in:0,1"  dc:"客服状态: 0-冻结 1-激活"`
}

type UpdateChatSupportStatusRes struct {
}
