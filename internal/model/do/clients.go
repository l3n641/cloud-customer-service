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
	CreateAt            *gtime.Time //
	Email               interface{} //
	DialCode            interface{} //
	Phone               interface{} //
	MerchantId          interface{} //
	Domain              interface{} //
	LastLoginTime       interface{} //
	LastSendMessageTime interface{} //
	Ip                  interface{} //
	BrowserFingerprint  interface{} //
	Area                interface{} //
	UserAgent           interface{} //
	Lang                interface{} //
	Status              interface{} //
	IsOnline            interface{} //
	Iso2                interface{} //
	StatusChangeTime    interface{} //
}
