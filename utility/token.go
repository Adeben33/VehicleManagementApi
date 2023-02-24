package utility

import (
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/config"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type CustomClaims struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(user model.User, secretKey string) string {
	claims := CustomClaims{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)).Unix(),
			IssuedAt:  jwt.NewNumericDate(time.Now()).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Panic(err)
	}
	return tokenString
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := config.GetConfig().Server.Secret
		return secretKey, nil
	})
	return token, err
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//	return claims, nil
	//} else {
	//	return nil, err
	//}
}
