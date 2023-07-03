package middleware

import (
	"github.com/gin-gonic/gin"
	"qyyh-go/database/table"
)

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user table.User
		token, err := c.Cookie("token")
		if err == nil {
			user = table.GetUserByToken(token)
			c.Set("user", user)
		}
	}
}
