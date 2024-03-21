// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"cloudCustomerService/api/admin"
)

type ControllerSession struct{}

func NewSession() admin.IAdminSession {
	return &ControllerSession{}
}
type ControllerAccount struct{}

func NewAccount() admin.IAdminAccount {
	return &ControllerAccount{}
}

type ControllerMerchant struct{}

func NewMerchant() admin.IAdminMerchant {
	return &ControllerMerchant{}
}

type ControllerChat_support struct{}
type ControllerChatSupport struct{}

func NewChatSupport() admin.IAdminChatSupport {
	return &ControllerChatSupport{}
}

