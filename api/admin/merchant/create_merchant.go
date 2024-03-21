package merchant

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateMerchantReq struct {
	g.Meta `path:"/merchant" method:"post" summary:"创建商户" tags:"管理员管理商户"`

	MerchantId   string `v:"required#请输入商户id" dc:"商户id"`
	MerchantName string `v:"required#请输入商户名称" dc:"商户名称"`
}

type CreateMerchantRes struct {
}
