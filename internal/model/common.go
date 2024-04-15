package model

type Pagination struct {
	Page int `json:"page"  dc:"分页号码，默认1"`
	Size int `json:"page_size" dc:"分页数量，最大100,默认20"`
}
