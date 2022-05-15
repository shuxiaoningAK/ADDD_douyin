package middleware

import (
	"ADDD_DOUYIN/util"
	"time"

	"github.com/gin-gonic/gin"
)

//实现自定义的token验证中间件
func TOKENPARSE() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		code := 1
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 0
		} else {
			_, claims, err := util.ParseToken(token)
			if err != nil {
				code = 0 //解析失败
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 0 //token过期
			}
		}
		if code != 1 {
			c.JSON(1, gin.H{
				"status": code,
				"msg":    "",
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
