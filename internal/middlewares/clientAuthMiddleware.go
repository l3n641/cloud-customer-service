package middlewares

import (
	"cloudCustomerService/api/client/session"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/oschwald/geoip2-golang"
	"net"
	"time"
)

var clientAuthService *jwt.GfJWTMiddleware

func ClientAuth() *jwt.GfJWTMiddleware {
	return clientAuthService
}

func init() {
	ctx := gctx.New()

	realm, _ := g.Cfg().Get(ctx, "jwt.realm", "app")
	key, _ := g.Cfg().Get(ctx, "jwt.secretKey", "secret key")
	identityKey := service.Client().GetIdentityKey()
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:         realm.String(),                                     // 用于展示中间件的名称
		Key:           key.Bytes(),                                        // token过期时间密钥
		Timeout:       time.Minute * 60 * 24 * 365,                        // token过期时间
		MaxRefresh:    time.Minute * 60 * 24 * 365,                        // token过期时间
		IdentityKey:   identityKey,                                        // 身份验证的key值
		TokenLookup:   "header: Authorization, query: token, cookie: jwt", // token检索模式，用于提取token-> Authorization
		TokenHeadName: "Bearer",                                           // token在请求头时的名称，默认值为Bearer;// 客户端在header中传入Authorization 对一个值是Bearer + 空格 + token
		TimeFunc:      time.Now,                                           // 测试或服务器在其他时区可设置该属性
		Authenticator: ClientAuthenticator,                                // 根据登录信息对用户进行身份验证的回调函数
		Unauthorized:  Unauthorized,                                       // 处理不进行授权的逻辑
		PayloadFunc:   PayloadFunc,                                        // 登录期间的回调的函数
		IdentityHandler: func(ctx context.Context) interface{} {
			claims := jwt.ExtractClaims(ctx)
			return claims[identityKey]
		},
	})
	clientAuthService = auth
}

func ClientAuthenticator(ctx context.Context) (interface{}, error) {
	var (
		request = g.RequestFromCtx(ctx)
		req     session.LoginReq
	)
	if err := request.Parse(&req); err != nil {
		return "", err
	}
	ipAddr := request.GetRemoteIp()
	var area, iso2 string
	ipDb, _ := g.Cfg().Get(ctx, "ipDb")

	country, err := Ip2Country(ipDb.String(), ipAddr)
	if country != nil {
		area = country.Country.Names["zh-CN"]
		iso2 = country.Country.IsoCode
	}
	loginOutput, err := service.Client().Login(ctx, &model.ClientLoginInput{
		LoginType:          req.LoginType,
		Email:              req.Email,
		Phone:              req.Phone,
		Domain:             req.Domain,
		MerchantId:         req.MerchantId,
		BrowserFingerprint: req.BrowserFingerprint,
		Ip:                 ipAddr,
		Area:               area,
		Iso2:               iso2,
		UserAgent:          request.Header.Get("User-Agent"),
		Lang:               request.Header.Get("Accept-Language"),
	},
	)
	if err != nil {
		return nil, err

	}
	return loginOutput.Data, nil
}

func Ip2Country(dbPath, ipStr string) (region *geoip2.Country, err error) {
	db, err := geoip2.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	ip := net.ParseIP(ipStr)
	record, err := db.Country(ip)

	return record, err

}
