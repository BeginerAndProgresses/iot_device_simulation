package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"iot_device_simulation/internal/consts"
	"iot_device_simulation/internal/service"
	"net/http"
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *iMiddleware {
	return &iMiddleware{}
}

type iMiddleware struct {
}

// Auth implements service.IMiddleware.
func (*iMiddleware) Auth(r *ghttp.Request) {
	// 过滤器
	path := []string{"/swagger", "/user/login", "/user/register"}
	if filter(r, path) {
		r.Middleware.Next()
		return
	}
	var res *ghttp.DefaultHandlerResponse
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		res = &ghttp.DefaultHandlerResponse{
			Code:    403,
			Message: "请登录后携带token访问",
		}
	} else {
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(consts.JwtTokenKey), nil
		})

		if err != nil || !token.Valid {
			res = &ghttp.DefaultHandlerResponse{
				Code:    403,
				Message: "token已失效，请重新登录",
			}
		}
	}
	if res != nil {
		r.Response.WriteJsonExit(res)
	}
	r.Middleware.Next()
}

func filter(r *ghttp.Request, path []string) bool {
	fmt.Println("请求路径：", r.URL.Path)
	for i := 0; i < len(path); i++ {
		if r.URL.Path == path[i] {
			return true
		}
	}
	return false
}

func (*iMiddleware) CORS(r *ghttp.Request) {
	fmt.Printf("请求：", r.URL)
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"localhost:5173"}
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
