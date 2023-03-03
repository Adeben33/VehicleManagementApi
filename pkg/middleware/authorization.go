package middleware

import (
	"fmt"
	"github.com/adeben33/vehicleParkingApi/utility"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		//	get the userToken from the cookie
		token, err := c.Request.Cookie("userToken")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No usercookie found"})
			return
		}
		//	Validate the userToken
		claims, errString := utility.ValidateToken(token.Value)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errString})
			return
		}
		//check for expires
		//if time.Now().Unix() > claim["ExpiresAt"].(time.Time).Unix() {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
		//	return
		//}

		//if !jwtToken.Valid {
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		//Claims
		//if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		//	if time.Now().Unix() > claims["ExpiresAt"].(time.Time).Unix() {
		//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
		//		return
		//	}
		//	userEmail := claims["email"]
		//	Role := claims["Role"]
		//	c.Set("userEmail", userEmail)
		//	c.Set("role", Role)
		//} else {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "claims not seen "})
		//	return
		//}
		fmt.Printf(claims["role"].(string))
		c.Next()
	}

}
