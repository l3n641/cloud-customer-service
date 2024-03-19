package middlewares

import (
	"cloudCustomerService/api/admin/session"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	ctx := gctx.New()

	realm, _ := g.Cfg().Get(ctx, "jwt.realm", "app")
	key, _ := g.Cfg().Get(ctx, "jwt.secretKey", "secret key")
	identityKey := service.Admin().GetIdentityKey()
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           realm.String(),                                     // 用于展示中间件的名称
		Key:             key.Bytes(),                                        // token过期时间密钥
		Timeout:         time.Minute * 60,                                   // token过期时间
		MaxRefresh:      time.Minute * 60,                                   // token过期时间
		IdentityKey:     identityKey,                                        // 身份验证的key值
		TokenLookup:     "header: Authorization, query: token, cookie: jwt", // token检索模式，用于提取token-> Authorization
		TokenHeadName:   "Bearer",                                           // token在请求头时的名称，默认值为Bearer;// 客户端在header中传入Authorization 对一个值是Bearer + 空格 + token
		TimeFunc:        time.Now,                                           // 测试或服务器在其他时区可设置该属性
		Authenticator:   Authenticator,                                      // 根据登录信息对用户进行身份验证的回调函数
		Unauthorized:    Unauthorized,                                       // 处理不进行授权的逻辑
		PayloadFunc:     PayloadFunc,                                        // 登录期间的回调的函数
		IdentityHandler: IdentityHandler,                                    // 解析并设置用户身份信息
	})
	authService = auth
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
// 此方法用于验证账户密码是否正确，基于自己的代码嵌套即可
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		request = g.RequestFromCtx(ctx)
		req     session.LoginReq
	)
	if err := request.Parse(&req); err != nil {
		return "", err
	}

	loginOutput, err := service.Admin().Login(ctx, &model.AdminLoginInput{
		Username: req.Username,
		Password: req.Password,
		Ip:       request.GetRemoteIp(),
	},
	)
	if err != nil {
		return nil, err

	}
	return loginOutput.Data, nil
}
