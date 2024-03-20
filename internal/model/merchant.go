package model

// MerchantCreateInput 创建商户
type MerchantCreateInput struct {
	MerchantId   string // 商户id
	MerchantName string // 商户名称
}

// MerchantSearchInput 搜索商户input
type MerchantSearchInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大100
	SearchFields
}

type SearchFields struct {
}

// SearchMerchantOutput 查询列表结果
type SearchMerchantOutput struct {
	List  []MerchantItem `json:"list" description:"列表"`
	Page  int            `json:"page" description:"分页码"`
	Size  int            `json:"size" description:"分页数量"`
	Total int            `json:"total" description:"数据总数"`
}

type MerchantItem struct {
	Id           int    `json:"id"`            //id
	MerchantId   string `json:"merchant_id"`   //商户id
	MerchantName string `json:"merchant_name"` //商户名称

}
