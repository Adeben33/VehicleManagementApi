package middleware

import (
	"github.com/adeben33/vehicleParkingApi/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		//	get the userToken from the cookie
		token, err := c.Request.Cookie("userToken")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		//	Validate the userToken
		jwtToken, err := utility.ValidateToken(token.Value)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if !jwtToken.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//Claims
		if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
			if time.Now().Unix() > claims["ExpiresAt"].(time.Time).Unix() {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
				return
			}
			userEmail := claims["email"]
			Role := claims["Role"]
			c.Set("userEmail", userEmail)
			c.Set("role", Role)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "claims not seen "})
			return
		}
		c.Next()
	}

}
