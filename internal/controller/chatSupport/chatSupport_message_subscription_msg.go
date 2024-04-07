package chatSupport

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
	"time"

	"cloudCustomerService/api/chatSupport/message"
)

func (c *ControllerMessage) SubscriptionMsg(ctx context.Context, req *message.SubscriptionMsgReq) (res *message.SubscriptionMsgRes, err error) {

	r := g.RequestFromCtx(ctx)

	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	redisClient := g.Redis(consts.RedisMsgQueueName)
	userId := gconv.Int(middlewares.ChatSupportAuth().GetIdentity(r.Context()))
	service.ChatSupport().UpdateSessionInfo(ctx, userId, consts.ChatSupportOnline, r.GetRemoteIp())
	r.Response.Writefln("event: connect\ndata: %s\n", "1")

	r.Response.Flush()
	for {
		select {
		case <-r.Context().Done():
			ctx := gctx.New()
			service.ChatSupport().UpdateSessionInfo(ctx, userId, consts.ChatSupportOutline, r.GetRemoteIp())
			return
		case <-time.After(5000 * time.Millisecond):
			key := consts.ChatSupportQueueKey + strconv.Itoa(userId)
			queueLen, err := redisClient.LLen(ctx, key)
			if queueLen < 1 || err != nil {
				continue
			}
			result, err := redisClient.LRange(ctx, key, 0, queueLen)
			if err != nil {
				continue
			}

			var messages []*model.ChatSupportMessageItem
			if err = result.Scan(&messages); err != nil {
				continue
			}
			redisClient.LTrim(ctx, key, queueLen, -1)

			b, _ := json.Marshal(messages)

			r.Response.Writefln("data: %s\n", string(b))

			r.Response.Flush()
		}
	}
}
