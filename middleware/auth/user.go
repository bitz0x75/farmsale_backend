package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"farmsale_backend/models/jwtmodel"
	jwt "github.com/dgrijalva/jwt-go"
)

// JwtVerify Middleware function
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := verifyTokenHelper(w, r)

		token, err := jwt.ParseWithClaims(tokenString, &jwtmodel.Token{}, func(token *jwt.Token) (interface{}, error) {
			return decodebs, nil
		})

		if _, ok := token.Claims.(*jwtmodel.Token); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "user", token)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "User not authorized"})
			fmt.Println(err)
			return
		}
	})
}
