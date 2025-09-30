package middleware

import (
	"fmt"
	"net/http"

	"example.com/net-http-class/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized access", "err": "No token in request"})
		return
	}

	claims, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized access", "err": err.Error(), "claims": claims})
		return
	}
	c.Set("userid", claims.User_ID)
	fmt.Println(claims.User_ID)
	c.Next()
}
