package models

import "github.com/golang-jwt/jwt/v4"

type AppClaimas struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
