package merchant

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DetailMerchantReq struct {
	g.Meta `path:"/merchant/{Id}" method:"get" summary:"获取商户详细信息" tags:"商户详细信息"`
	Id     int `json:"id" dc:"商户记录id"`
}

type DetailMerchantRes struct {
	Id           int    `json:"id"        `                    // id
	MerchantId   string `json:"merchant_id"            `       // 商户id
	MerchantName string `json:"merchant_name"                ` //商户名称
}
