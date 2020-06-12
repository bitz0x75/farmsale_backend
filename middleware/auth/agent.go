package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/maxwellgithinji/farmsale_backend/models/jwtmodel"
)

func AgentVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := verifyTokenHelper(w, r)

		token, err := jwt.ParseWithClaims(tokenString, &jwtmodel.Token{}, func(token *jwt.Token) (interface{}, error) {
			return decodebs, nil
		})

		if claims, ok := token.Claims.(*jwtmodel.Token); ok && token.Valid {
			if claims.Userclass != "agent" {
				//Give the super admin the green light to access manager toute
				if claims.Isadmin {
					ctx := context.WithValue(r.Context(), "user", token)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				//Give the manager oversight over their agents
				if claims.Userclass == "manager" {
					ctx := context.WithValue(r.Context(), "user", token)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(Exception{Message: "User not authorized, Admin or manager or agent only"})
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
