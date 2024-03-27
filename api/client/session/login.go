package session

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type LoginReq struct {
	g.Meta `path:"/session" method:"post" summary:"客户登录登录" tags:"客户管理"`

	LoginType          int    `v:"required|in:1,2" json:"login_type" dc:"登录类型:1-邮箱 2-电话号码"`
	Email              string `v:"required-if:login_type,1" dc:"邮箱"`
	Phone              string `v:"required-if:login_type,2" dc:"手机号"`
	Domain             string `v:"required|domain"  dc:"域名"`
	MerchantId         string `v:"required"  json:"merchant_id" dc:"商户号"`
	BrowserFingerprint string `json:"browser_fingerprint" dc:"浏览器指纹"`
}

type LoginRes struct {
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}
