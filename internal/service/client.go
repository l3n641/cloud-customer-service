// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"context"
)

type (
	IClient interface {
		GetClientById(ctx context.Context, id int) (*entity.Clients, error)
		GetIdentityKey() string
		Login(ctx context.Context, in *model.ClientLoginInput) (*model.ClientLoginOutput, error)
		SetAccountStatus(ctx context.Context, clientId, status int) (int64, error)
	}
)

var (
	localClient IClient
)

func Client() IClient {
	if localClient == nil {
		panic("implement not found for interface IClient, forgot register?")
	}
	return localClient
}

func RegisterClient(i IClient) {
	localClient = i
}
