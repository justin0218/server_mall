package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server_mall/pkg/jwt"
	"server_mall/pkg/resp"
	"server_mall/store"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if new(store.Config).Get().Runmode == "debug" {
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
