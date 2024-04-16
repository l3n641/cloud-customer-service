package client

import (
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/client/message"
)

func (c *ControllerMessage) GetWelcomeMessage(ctx context.Context, req *message.GetWelcomeMessageReq) (res *message.GetWelcomeMessageRes, err error) {
	data, err := service.MessageTemplate().GetWelcomeMessage(ctx)
	return &message.GetWelcomeMessageRes{
		Item: data,
	}, err

}
