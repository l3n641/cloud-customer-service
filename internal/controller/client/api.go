package client

import (
	"cloudCustomerService/api/client/message"
	"cloudCustomerService/api/client/session"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ControllerApi struct{}

func NewApi() *ControllerApi {
	return &ControllerApi{}
}

func (c *ControllerApi) Router(r *ghttp.Request) {
	action := r.Get("action")
	ctx := r.Context()
	var err error
	var res interface{}

	switch action.String() {
	case "login":
		var data *session.LoginReq
		r.Parse(&data)
		res, err = NewSession().Login(ctx, data)

	case "message-list":
		var data *message.GetMessageListReq
		if err = r.Parse(&data); err != nil {
			break
		}

		res, err = NewMessage().GetMessageList(ctx, data)

	case "message-quantity":
		var data *message.GetUnreadMessageQuantityReq
		if err = r.Parse(&data); err != nil {
			break
		}

		res, err = NewMessage().GetUnreadMessageQuantity(ctx, data)

	case "send-message":
		var data *message.SendMessageReq
		err = r.Parse(&data)
		if err = r.Parse(&data); err != nil {
			break
		}
		res, err = NewMessage().SendMessage(ctx, data)
	}

	r.SetError(err)

	if err == nil {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    0,
			Message: "",
			Data:    res,
		})
	}

}
