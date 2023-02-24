package utility

import (
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
