package client

import "github.com/gogf/gf/v2/frame/g"

type ManageClientAccountBanReq struct {
	g.Meta `path:"/client-ban" method:"put" summary:"设置账号状态" tags:"客服模块"`

	ClientId int `v:"required" json:"client_id" dc:"客人id"`
	Status   int `v:"required|in:0,1" json:"status" dc:"账号状态 0-解除封号 1-设置封号"`
}

type ManageClientAccountBanRes struct {
}
