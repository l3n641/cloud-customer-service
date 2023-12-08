package admin

import (
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"

	"cloudCustomerService/api/admin/customerServiceAgent"
)

func (c *ControllerCustomerServiceAgent) CreateCustomerServiceAgents(ctx context.Context, req *customerServiceAgent.CreateCustomerServiceAgentsReq) (res *customerServiceAgent.CreateCustomerServiceAgentsRes, err error) {
	err = service.CustomerServiceAgent().Crete(ctx, &model.CreateCustomerServiceAgentInput{
		Email:    req.Email,
		Password: req.Password,
		NickName: req.Nickname,
	})
	if err != nil {
		return nil, err
	}

	return &customerServiceAgent.CreateCustomerServiceAgentsRes{}, nil

}
