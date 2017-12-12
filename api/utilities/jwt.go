package utilities

import (
	"github.com/dgrijalva/jwt-go"
)

// Token secret for authorization through JWT
var GP_TOKEN_SECRET = []byte(GetVar("GP_TOKEN_SECRET", DEFAULT_TOKEN_SECRET))

// ExtractClaims will take the claim out of an authorization header.
// It takes in the JWT tokenString to read the claims from.
// It returns a map of string to interface for the claims.
func ExtractClaims(tokenString string) map[string]interface{} {
	if tokenString == "" {
		return nil
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return GP_TOKEN_SECRET, nil
	})
	return token.Claims.(jwt.MapClaims)
}
