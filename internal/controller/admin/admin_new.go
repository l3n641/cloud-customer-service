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

type ControllerCustomerServiceAgent struct{}

func NewCustomerServiceAgent() admin.IAdminCustomerServiceAgent {
	return &ControllerCustomerServiceAgent{}
}

