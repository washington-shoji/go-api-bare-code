package midleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/washington-shoji/gobare/configs"
	"github.com/washington-shoji/gobare/databases"
	"github.com/washington-shoji/gobare/helpers"
)

func JWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}

		userID, err := helpers.GetID(r)
		if err != nil {
			permissionDenied(w)
			return
		}
		account, err := databases.GetAccountByID(userID)
		if err != nil {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if account.UserName != claims["userName"] {
			permissionDenied(w)
			return
		}

		if err != nil {
			permissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}
}

func validateJWT(tokenStr string) (*jwt.Token, error) {
	secret := configs.EnvConfig("JWT_SECRET")

	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

}

func permissionDenied(w http.ResponseWriter) {
	helpers.WriteJson(w, http.StatusForbidden, databases.ApiError{Error: "permission denied"})
}
