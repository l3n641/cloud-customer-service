package api

type PaginationReq struct {
	Page     int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	PageSize int `json:"page_size" in:"query" d:"20" v:"max:100#分页数量最大100条" dc:"分页数量，最大100,默认20"`
}

type PaginationRes struct {
	Item     interface{} `json:"item" dc:"列表数据"`
	Total    int         `json:"total" dc:"总数"`
	Page     int         `json:"page" dc:"分页号码"`
	PageSize int         `json:"page_size" dc:"分页数量"`
}
