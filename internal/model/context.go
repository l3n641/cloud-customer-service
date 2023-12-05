package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IContext interface {
	GetData() g.Map
	GetUserId() int
	SetUserId(int)
	SetData(g.Map)
}

// AdminContext 请求上下文结构
type AdminContext struct {
	UserId int
	Data   g.Map // 自定KV变量，业务模块根据需要设置，不固定
}

func (m *AdminContext) GetData() g.Map {
	return m.Data
}

func (m *AdminContext) SetData(data g.Map) {
	m.Data = data
}

func (m *AdminContext) GetUserId() int {
	return m.UserId
}

func (m *AdminContext) SetUserId(userId int) {
	m.UserId = userId
}
