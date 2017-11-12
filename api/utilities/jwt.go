package utilities

import (
	"github.com/dgrijalva/jwt-go"
)

func ExtractClaims(tokenString string) map[string]interface{} {
	if tokenString == "" {
		return nil
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return GP_TOKEN_SECRET, nil
	})
	return token.Claims.(jwt.MapClaims)
}
