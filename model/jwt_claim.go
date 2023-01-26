package model

import "github.com/golang-jwt/jwt/v4"

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
