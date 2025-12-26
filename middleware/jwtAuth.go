package middleware

import (
	"go-learn/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1、获取 token
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "缺少认证 token"})
			c.Abort() // 终止后续处理
			return
		}

		// 2、验证 token
		_, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort() // 终止后续处理
			return
		}

		// 3、处理后续请求
		c.Next()
	}
}
