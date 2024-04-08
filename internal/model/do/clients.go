// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Clients is the golang structure of table clients for DAO operations like Where/Data.
type Clients struct {
	g.Meta              `orm:"table:clients, do:true"`
	Id                  interface{} //
	CreateAt            *gtime.Time // 创建时间
	Email               interface{} // 客人邮箱
	DialCode            interface{} // 电话区号
	Phone               interface{} // 电话号码
	MerchantId          interface{} // 商户id
	Domain              interface{} // 域名
	LastLoginTime       *gtime.Time // 最后登录时间
	LastSendMessageTime interface{} // 最后发送消息时间
	Ip                  interface{} // 登录ip
	BrowserFingerprint  interface{} // 浏览器指纹
	Area                interface{} // 国家
	UserAgent           interface{} // 浏览器UA
	Lang                interface{} // 浏览器语言
	Status              interface{} // 账号状态 0-正常 1-被冻结
	IsOnline            interface{} // 在线状态 0-不在线 1-在线
	Iso2                interface{} // 国家区号
	StatusChangeTime    *gtime.Time // 账号状态修改时间
}
