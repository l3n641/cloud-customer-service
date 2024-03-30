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
