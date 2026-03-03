package utils

import (
	"time"

	"github.com/farhandwiyan/project-himaose/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// generate token
// generate refresh token
func GenerateToken(userID int64, role,username string, publicID uuid.UUID) (string, error) {
	secret := config.AppConfig.JWTSecret
	duration, _ := time.ParseDuration(config.AppConfig.JWTExpired)

	claims := jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"pub_id": publicID,
		"username":username,
		"exp":time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(userID int64) (string, error) {
	secret := config.AppConfig.JWTSecret
	duration, _ := time.ParseDuration(config.AppConfig.RefreshToken)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(secret))
}
