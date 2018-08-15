package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang-rest/utils"
	"net/http"
	"strings"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		// extract authorization header
		headerString := c.GetHeader("Authorization")

		// if there is not authorization header, throw error.
		if len(headerString) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Missing authorization header.",
			})

			c.Abort()

			// if there is not valid format token, throw error.
		} else if headerPairs := strings.Split(headerString, " "); len(headerPairs) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header format is invalid.",
			})
			c.Abort()

			// if jwt token does not exist, throw error.
		} else if claims, err := utils.CheckJwtToken(headerPairs[1]); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
		} else {
			c.Set("claims", claims)
			c.Next()
		}
	}
}
