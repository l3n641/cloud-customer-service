package merchant

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateMerchantReq struct {
	g.Meta `path:"/merchant/{Id}" method:"put" summary:"更新商户" tags:"管理员管理商户"`

	Id           int    `json:"id" dc:"商户记录id"`
	MerchantId   string `v:"required#请输入商户id" dc:"商户id"`
	MerchantName string `v:"required#请输入商户名称" dc:"商户名称"`
}

type UpdateMerchantRes struct {
}
