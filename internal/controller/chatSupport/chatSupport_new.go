// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chatSupport

import (
	"cloudCustomerService/api/chatSupport"
)

type ControllerSession struct{}

func NewSession() chatSupport.IChatSupportSession {
	return &ControllerSession{}
}
type ControllerAccount struct{}

func NewAccount() chatSupport.IChatSupportAccount {
	return &ControllerAccount{}
}

type ControllerTicket struct{}

func NewTicket() chatSupport.IChatSupportTicket {
	return &ControllerTicket{}
}

type ControllerMessage struct{}

func NewMessage() chatSupport.IChatSupportMessage {
	return &ControllerMessage{}
}

