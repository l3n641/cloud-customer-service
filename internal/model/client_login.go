package model

import "github.com/gogf/gf/v2/frame/g"

// ClientLoginInput 客户登录
type ClientLoginInput struct {
	LoginType          int
	Email              string
	Phone              string
	Domain             string
	MerchantId         string
	Ip                 string
	BrowserFingerprint string
	Area               string
	Iso2               string
	UserAgent          string
	Lang               string
}

type ClientLoginOutput struct {
	Data g.Map
}
