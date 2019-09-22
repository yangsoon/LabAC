package jwt

import (
	"net/http"
	"time"

	"labac/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = 1
		token := c.Query("token")
		if token == "" {
			code = 0
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = 0
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 0
			}
		}

		if code != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "非法操作",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
