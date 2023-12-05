// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"cloudCustomerService/internal/model"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IBizCtx interface {
		// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
		Init(r *ghttp.Request, customCtx model.IContext)
		// GetContext 获得上下文变量，如果没有设置，那么返回nil
		GetContext(ctx context.Context) model.IContext
		// SetData 将上下文信息设置到上下文请求中，注意是完整覆盖
		SetData(ctx context.Context, data g.Map)
		// SetUserId 将上下文信息设置到上下文请求中，注意是完整覆盖
		SetUserId(ctx context.Context, userId int)
	}
)

var (
	localBizCtx IBizCtx
)

func BizCtx() IBizCtx {
	if localBizCtx == nil {
		panic("implement not found for interface IBizCtx, forgot register?")
	}
	return localBizCtx
}

func RegisterBizCtx(i IBizCtx) {
	localBizCtx = i
}
