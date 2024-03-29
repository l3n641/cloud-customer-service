// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Clients is the golang structure for table clients.
type Clients struct {
	Id                  int         `json:"id"                  ` //
	CreateAt            *gtime.Time `json:"createAt"            ` //
	Email               string      `json:"email"               ` //
	DialCode            string      `json:"dialCode"            ` //
	Phone               string      `json:"phone"               ` //
	MerchantId          string      `json:"merchantId"          ` //
	Domain              string      `json:"domain"              ` //
	LastLoginTime       *gtime.Time `json:"lastLoginTime"       ` //
	LastSendMessageTime int         `json:"lastSendMessageTime" ` //
	Ip                  string      `json:"ip"                  ` //
	BrowserFingerprint  string      `json:"browserFingerprint"  ` //
	Area                string      `json:"area"                ` //
	UserAgent           string      `json:"userAgent"           ` //
	Lang                string      `json:"lang"                ` //
	Status              int         `json:"status"              ` //
	IsOnline            int         `json:"isOnline"            ` //
	Iso2                string      `json:"iso2"                ` //
	StatusChangeTime    int         `json:"statusChangeTime"    ` //
}
