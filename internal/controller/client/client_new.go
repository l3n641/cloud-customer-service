// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package client

import (
	"cloudCustomerService/api/client"
)

type ControllerSession struct{}

func NewSession() client.IClientSession {
	return &ControllerSession{}
}

type ControllerMessage struct{}

func NewMessage() client.IClientMessage {
	return &ControllerMessage{}
}
type ControllerChatSupport struct{}

func NewChatSupport() client.IClientChatSupport {
	return &ControllerChatSupport{}
}

