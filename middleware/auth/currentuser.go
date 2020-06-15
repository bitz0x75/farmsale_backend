package auth

import (
	"context"
	"encoding/json"
	"github.com/maxwellgithinji/farmsale_backend/models/jwtmodel"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// CurrentUserVerify ensures only the profile owner can edit their details
func CurrentUserVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := verifyTokenHelper(w, r)

		token, err := jwt.ParseWithClaims(tokenString, &jwtmodel.Token{}, func(token *jwt.Token) (interface{}, error) {
			return decodebs, nil
		})
		params := mux.Vars(r)

		if claims, ok := token.Claims.(*jwtmodel.Token); ok && token.Valid {
			if claims.ID.Hex() != params["id"] {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(Exception{Message: "User not authorized"})
				fmt.Println(err)
				return
			}
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
