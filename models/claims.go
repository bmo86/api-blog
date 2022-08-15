package models

import "github.com/golang-jwt/jwt"

type AppClaimas struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
