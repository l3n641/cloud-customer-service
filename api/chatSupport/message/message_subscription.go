package message

import "github.com/gogf/gf/v2/frame/g"

type SubscriptionMsgReq struct {
	g.Meta `path:"/sub-msg" method:"get" summary:"通过sse订阅消息" tags:"客服模块"`
}

type SubscriptionMsgRes struct {
}
