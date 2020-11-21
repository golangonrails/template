package jwtutil

import (
	"app/config"
	"unsafe"

	"github.com/dgrijalva/jwt-go"
)

// Secret of jwt
var Secret = func() []byte {
	return *(*[]byte)(unsafe.Pointer(&config.Settings().JWT.Secret))
}

func keyFromToken(token *jwt.Token) (interface{}, error) {
	return Secret(), nil
}

// New JWT from claims
func New(claims jwt.Claims) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key, err := keyFromToken(token)
	if err != nil {
		return
	}
	return token.SignedString(key)
}

//Parse JWT from signedToken, claimsOut save claims value
func Parse(signedToken string, claimsOut jwt.Claims) (token *jwt.Token, err error) {
	return jwt.ParseWithClaims(signedToken, claimsOut, keyFromToken)
}
