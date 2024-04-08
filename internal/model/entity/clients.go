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
	CreateAt            *gtime.Time `json:"createAt"            ` // 创建时间
	Email               string      `json:"email"               ` // 客人邮箱
	DialCode            string      `json:"dialCode"            ` // 电话区号
	Phone               string      `json:"phone"               ` // 电话号码
	MerchantId          string      `json:"merchantId"          ` // 商户id
	Domain              string      `json:"domain"              ` // 域名
	LastLoginTime       *gtime.Time `json:"lastLoginTime"       ` // 最后登录时间
	LastSendMessageTime int         `json:"lastSendMessageTime" ` // 最后发送消息时间
	Ip                  string      `json:"ip"                  ` // 登录ip
	BrowserFingerprint  string      `json:"browserFingerprint"  ` // 浏览器指纹
	Area                string      `json:"area"                ` // 国家
	UserAgent           string      `json:"userAgent"           ` // 浏览器UA
	Lang                string      `json:"lang"                ` // 浏览器语言
	Status              int         `json:"status"              ` // 账号状态 0-正常 1-被冻结
	IsOnline            int         `json:"isOnline"            ` // 在线状态 0-不在线 1-在线
	Iso2                string      `json:"iso2"                ` // 国家区号
	StatusChangeTime    *gtime.Time `json:"statusChangeTime"    ` // 账号状态修改时间
}
