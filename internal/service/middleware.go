// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IAdminAuthMiddleware interface {
		// AdminCtx 自定义上下文对象
		Ctx(r *ghttp.Request)
		Auth(r *ghttp.Request)
		GetJWTMiddleware(ctx context.Context) *jwt.GfJWTMiddleware
	}
)

var (
	localAdminAuthMiddleware IAdminAuthMiddleware
)

func AdminAuthMiddleware() IAdminAuthMiddleware {
	if localAdminAuthMiddleware == nil {
		panic("implement not found for interface IAdminAuthMiddleware, forgot register?")
	}
	return localAdminAuthMiddleware
}

func RegisterAdminAuthMiddleware(i IAdminAuthMiddleware) {
	localAdminAuthMiddleware = i
}
