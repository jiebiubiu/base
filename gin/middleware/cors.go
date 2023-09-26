package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	allowOrigin = map[string]struct{}{
//		"*":                     {},
//		"http://localhost:5005": {},
//	}
func Cors(allowOrigin map[string]struct{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义一个origin的map，只有在字典中的key才允许跨域请求

		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			if _, ok := allowOrigin["*"]; !ok {
				if _, ok := allowOrigin[origin]; !ok {
					c.AbortWithStatus(http.StatusNoContent)
					return
				}
			}

			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			// c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type,Content-Length,X-CSRF-Token,Token,X-Token,session,X_Requested_With,Accept,Origin,Host,Connection,Accept-Encoding,Accept-Language,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Pragma,Sec-Ch-Ua-Platform,Sec-Ch-Ua-Mobile,Sec-Ch-Ua,Referer")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true") //  跨域请求是否需要带cookie信息 默认设置为true
		}

		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
