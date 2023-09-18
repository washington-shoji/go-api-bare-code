package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/washington-shoji/gobare/configs"
)

func CreateJWT(cred string) (string, error) {
	// Create the Claims
	claims := &jwt.MapClaims{
		"expiredAt": time.Now().Add(time.Minute * 1).Unix(),
		"userName":  cred,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := configs.EnvConfig("JWT_SECRET")

	return token.SignedString([]byte(secret))
}
