package middleware

import (
	"Go-RestfulAPI/handler"
	"Go-RestfulAPI/pkg/errno"
	"Go-RestfulAPI/pkg/token"
	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2020-08-26 20:19
* @Description:
**/

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
