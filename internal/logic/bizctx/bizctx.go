package bizctx

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sBizCtx struct{}

func init() {
	service.RegisterBizCtx(New())
}

func New() *sBizCtx {
	return &sBizCtx{}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sBizCtx) Init(r *ghttp.Request, customCtx model.IContext) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// GetContext 获得上下文变量，如果没有设置，那么返回nil
func (s *sBizCtx) GetContext(ctx context.Context) model.IContext {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(model.IContext); ok {
		return localCtx
	}
	return nil
}

// SetData 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sBizCtx) SetData(ctx context.Context, data g.Map) {
	localCtx := s.GetContext(ctx)
	if localCtx != nil {
		localCtx.SetData(data)
	}
}

// SetUserId 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sBizCtx) SetUserId(ctx context.Context, userId int) {
	localCtx := s.GetContext(ctx)
	if localCtx != nil {
		localCtx.SetUserId(userId)
	}
}
