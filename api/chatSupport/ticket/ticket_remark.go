package ticket

import "github.com/gogf/gf/v2/frame/g"

type UpdateTicketRemarkReq struct {
	g.Meta `path:"/ticket-remark/{Id}" method:"put" summary:"修改工单备注" tags:"客服管理"`

	Id     int    `json:"id" dc:"工单id"`
	Remark string `json:"remark" dc:"备注"`
}

type UpdateTicketRemarkRes struct {
}
