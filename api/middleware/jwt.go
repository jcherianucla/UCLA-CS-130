package middleware

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.Handler {
	JWTMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return utilities.GP_TOKEN_SECRET, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return JWTMiddleware.Handler(next)
}
