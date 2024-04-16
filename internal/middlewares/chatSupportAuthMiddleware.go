package middlewares

import (
	"cloudCustomerService/api/chatSupport/session"
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

var chatSupportAuthService *jwt.GfJWTMiddleware

func ChatSupportAuth() *jwt.GfJWTMiddleware {
	return chatSupportAuthService
}

func init() {
	ctx := gctx.New()

	realm, _ := g.Cfg().Get(ctx, "jwt.realm", "app")
	key, _ := g.Cfg().Get(ctx, "jwt.secretKey", "secret key")
	identityKey := service.ChatSupport().GetIdentityKey()
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:         realm.String(),                                     // 用于展示中间件的名称
		Key:           key.Bytes(),                                        // token过期时间密钥
		Timeout:       time.Minute * 60 * 24 * 365,                        // token过期时间
		MaxRefresh:    time.Minute * 60 * 24 * 365,                        // token过期时间
		IdentityKey:   identityKey,                                        // 身份验证的key值
		TokenLookup:   "header: Authorization, query: token, cookie: jwt", // token检索模式，用于提取token-> Authorization
		TokenHeadName: "Bearer",                                           // token在请求头时的名称，默认值为Bearer;// 客户端在header中传入Authorization 对一个值是Bearer + 空格 + token
		TimeFunc:      time.Now,                                           // 测试或服务器在其他时区可设置该属性
		Authenticator: ChatSupportAuthenticator,                           // 根据登录信息对用户进行身份验证的回调函数
		Unauthorized:  Unauthorized,                                       // 处理不进行授权的逻辑
		PayloadFunc:   PayloadFunc,                                        // 登录期间的回调的函数
		IdentityHandler: func(ctx context.Context) interface{} {
			claims := jwt.ExtractClaims(ctx)
			return claims[identityKey]
		},
		Authorizator: func(clientId interface{}, ctx context.Context) bool {
			user, err := service.ChatSupport().GetDetailById(ctx, int(clientId.(float64)))
			if err != nil || user == nil {
				return false
			}
			if user.Status == consts.ChatSupportInactive {
				return false
			}
			return true
		},
	})
	chatSupportAuthService = auth
}

func ChatSupportAuthenticator(ctx context.Context) (interface{}, error) {
	var (
		request = g.RequestFromCtx(ctx)
		req     session.LoginReq
	)
	if err := request.Parse(&req); err != nil {
		return "", err
	}

	loginOutput, err := service.ChatSupport().Login(ctx, &model.ChatSupportLoginInput{
		Email:    req.Email,
		Password: req.Password,
		Ip:       request.GetRemoteIp(),
	},
	)
	if err != nil {
		return nil, err

	}
	return loginOutput.Data, nil
}
