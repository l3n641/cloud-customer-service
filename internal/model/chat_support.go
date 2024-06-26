package model

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatSupportCreateInput 创建客服输入
type ChatSupportCreateInput struct {
	Email         string // 账号
	Password      string // 密码(明文)
	NickName      string // 昵称
	MerchantGroup []int  //要关联的商户

}

// ChatSupportCreateOutput 创建客服输出
type ChatSupportCreateOutput struct {
}

// ChatSupportSearchInput 搜索客服input
type ChatSupportSearchInput struct {
	Page    int    // 分页号码
	Size    int    // 分页数量，最大100
	OrderBy string //排序
	SearchChatSupportFields
}

type SearchChatSupportFields struct {
	Email    string
	Status   gvar.Var
	IsOnline gvar.Var
}

// ChatSupportSearchOutput 查询列表结果
type ChatSupportSearchOutput struct {
	List  []ChatSupportItem `json:"list" description:"列表"`
	Page  int               `json:"page" description:"分页码"`
	Size  int               `json:"size" description:"分页数量"`
	Total int               `json:"total" description:"数据总数"`
}

type ChatSupportItem struct {
	Id          int         `json:"id"`
	CreateAt    *gtime.Time `json:"create_at"`
	Email       string      `json:"email"`
	LastLoginAt *gtime.Time `json:"last_login_at"`
	LastLoginIp string      `json:"last_login_ip"`
	IsOnline    int         `json:"is_online"`
	Status      int         `json:"status"`
}

type ChatSupportDetail struct {
	ChatSupportItem
	MerchantGroup []*ChatSupportMerchant
}

type ChatSupportMerchant struct {
	MerchantId int `json:"merchant_id"`
}

// ChatSupportUpdateInput 更新商户信息
type ChatSupportUpdateInput struct {
	Id            int
	Password      string // 密码(明文)
	NickName      string // 昵称
	MerchantGroup []int  //要关联的商户
}

// ChatSupportLoginInput 用户登录
type ChatSupportLoginInput struct {
	Email    string // 用户名
	Password string // 密码(明文)
	Ip       string //
}

type ChatSupportLoginOutput struct {
	Data g.Map
}

type IsOnlineChatSupport struct {
	IsOnline int                `json:"is_online" description:"客服是否在线 0-不在线 1-在线"`
	Messages []MessageTemplates `json:"messages" description:"消息模板列表"`
}

type MessageTemplates struct {
	Lang    string `json:"lang" description:"语种"`
	Content string `json:"content" description:"内容"`
}
