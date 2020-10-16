package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server_mall/configs"
	"server_mall/pkg/jwt"
	"server_mall/pkg/resp"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if configs.Dft.Get().Runmode == "debug" {
			c.Next()
			return
		}
		token := c.Query("token")
		uid, err := jwt.VerifyToken(token)
		if err != nil {
			resp.RespCode(c, http.StatusUnauthorized, "未授权")
			c.Abort()
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}
