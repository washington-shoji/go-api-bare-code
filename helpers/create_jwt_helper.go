package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/washington-shoji/gobare/configs"
)

func CreateJWT(cred string) (string, error) {
	// Create the Claims
	claims := &jwt.MapClaims{
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
		"userName": cred,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := configs.EnvConfig("JWT_SECRET")

	return token.SignedString([]byte(secret))
}
