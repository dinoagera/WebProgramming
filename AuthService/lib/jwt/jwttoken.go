package lib

import (
	"authservice/internal/config"
	"authservice/internal/domain/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user models.User) (string, error) {
	cfg := config.GetConfig()
	secretKey := cfg.SecretKey
	claims := jwt.MapClaims{
		"uid":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(cfg.TokenTTL).Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil

}
