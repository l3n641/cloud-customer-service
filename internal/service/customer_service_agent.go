// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"cloudCustomerService/internal/model"
	"context"
)

type (
	ICustomerServiceAgent interface {
		Crete(ctx context.Context, in *model.CreateCustomerServiceAgentInput) (err error)
	}
)

var (
	localCustomerServiceAgent ICustomerServiceAgent
)

func CustomerServiceAgent() ICustomerServiceAgent {
	if localCustomerServiceAgent == nil {
		panic("implement not found for interface ICustomerServiceAgent, forgot register?")
	}
	return localCustomerServiceAgent
}

func RegisterCustomerServiceAgent(i ICustomerServiceAgent) {
	localCustomerServiceAgent = i
}
