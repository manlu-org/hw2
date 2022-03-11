package middleware

import (
	"backend-learning/hw2/models"
	"backend-learning/hw2/pkg/request"
	"backend-learning/hw2/pkg/response"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

var GinJWTMiddleware *jwt.GinJWTMiddleware

type User struct {
	UserName string
}

func InitJWT() {
	//global.Log.Debug("Hello")
	GinJWTMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "hw2",                          // jwt标识
		Key:             []byte("backend-learning-hw2"), // 服务端密钥
		IdentityKey:     jwt.IdentityKey,
		Timeout:         time.Hour * 24,                                     // token过期时间
		MaxRefresh:      time.Hour * 24,                                     // token最大刷新时间(RefreshToken过期时间=Timeout+MaxRefresh)
		PayloadFunc:     payloadFunc,                                        // 有效载荷处理
		IdentityHandler: identityHandler,                                    // 解析Claims
		Authenticator:   login,                                              // 校验token的正确性, 处理登录逻辑
		Authorizator:    authorizator,                                       // 用户登录校验成功处理
		Unauthorized:    unauthorized,                                       // 用户登录校验失败处理
		LoginResponse:   loginResponse,                                      // 登录成功后的响应
		LogoutResponse:  logoutResponse,                                     // 登出后的响应
		RefreshResponse: refreshResponse,                                    // 刷新token后的响应
		TokenLookup:     "header: Authorization, query: token, cookie: jwt", // 自动在这几个地方寻找请求中的token
		TokenHeadName:   "Bearer",                                           // header名称
		TimeFunc:        time.Now,
	})
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*User); ok {
		return jwt.MapClaims{
			"user": v.UserName,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &User{
		UserName: claims["user"].(string),
	}
}

func login(c *gin.Context) (interface{}, error) {
	var params request.SignInRequest
	_ = c.ShouldBind(&params)
	if params.Username == "admin" && params.Password == "123456" {
		return &User{UserName: "admin"}, nil
	} else {
		return nil, jwt.ErrFailedAuthentication
	}
}

func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*User); ok && v.UserName == "admin" {
		return true
	}

	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	response.FailWithMessage(message)
}

func loginResponse(c *gin.Context, code int, token string, expires time.Time) {
	response.SuccessWithData(map[string]interface{}{
		"token": token,
		"expires": models.LocalTime{
			Time: expires,
		},
	})
}

func logoutResponse(c *gin.Context, code int) {
	response.Success()
}

func refreshResponse(c *gin.Context, code int, token string, expires time.Time) {
	response.SuccessWithData(map[string]interface{}{
		"token": token,
		"expires": models.LocalTime{
			Time: expires,
		},
	})
}
