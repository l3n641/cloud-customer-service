package middleware

import (
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type sAdminAuthMiddleware struct {
	identityKey string
}

func init() {
	service.RegisterAdminAuthMiddleware(NewAdminAuthMiddleware())
}

func NewAdminAuthMiddleware() *sAdminAuthMiddleware {
	return &sAdminAuthMiddleware{}
}

// Ctx 自定义上下文对象
func (s *sAdminAuthMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.AdminContext{
		Data: make(g.Map),
	}
	service.BizCtx().Init(r, customCtx)

	// 执行下一步请求逻辑
	r.Middleware.Next()
}

func (s *sAdminAuthMiddleware) Auth(r *ghttp.Request) {
	ctx := r.GetCtx()
	jwtMiddleware := s.GetJWTMiddleware(ctx)
	jwtMiddleware.MiddlewareFunc()(r)

	userId := gconv.Int(jwtMiddleware.GetIdentity(ctx))
	admin, err := service.Admin().GetUserInfo(ctx, userId)
	if err != nil {
		return
	}

	service.BizCtx().SetUserId(ctx, admin.Id)

	r.Middleware.Next()
}

func (s *sAdminAuthMiddleware) GetJWTMiddleware(ctx context.Context) *jwt.GfJWTMiddleware {

	realm, _ := g.Cfg().Get(ctx, "jwt.realm", "app")
	key, _ := g.Cfg().Get(ctx, "jwt.secretKey", "secret key")
	identityKey := service.Admin().GetIdentityKey()
	s.identityKey = identityKey

	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           realm.String(),
		Key:             key.Bytes(),
		Timeout:         time.Minute * 60 * 24,
		MaxRefresh:      time.Minute * 60 * 24,
		IdentityKey:     identityKey,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   s.authenticator,
		PayloadFunc:     s.payloadFunc,
		IdentityHandler: s.identityHandler,
	})
	return auth
}

func (s *sAdminAuthMiddleware) authenticator(ctx context.Context) (interface{}, error) {

	context := service.BizCtx().GetContext(ctx)
	data := context.GetData()
	if data != nil {
		return data, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func (s *sAdminAuthMiddleware) identityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[s.identityKey]
}

func (s *sAdminAuthMiddleware) payloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}
