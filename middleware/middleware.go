package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const passcode = "5672139asdaw"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenHeader := c.GetHeader("Token") //Grab the token from the header
		if tokenHeader == "" { //Token is missing, returns with error code 401 Unauthorized
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
				"error": "Missing auth token",
			})
			return
		}

		if tokenHeader != passcode { //Token is invalid
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H {
				"error": "Token is not valid",
			})
			return
		}

		c.Next()//proceed in the middleware chain!

	}

}
