package means

import (
	"net/http"
	"rest-websockets/models"
	"rest-websockets/server"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func Token(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaimas{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})

	return token, err
}
