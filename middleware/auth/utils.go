package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"farmsale_backend/utils"
)

var (
	decodebs = []byte(os.Getenv("TOKEN_SECRET"))
)

//Exception struct
type Exception utils.Exception

func verifyTokenHelper(w http.ResponseWriter, r *http.Request) string {
	var tokenString = r.Header.Get("Authorization") //Grab the token from the header
	tokenString = strings.Split(tokenString, "Bearer ")[1]
	if tokenString == "" {
		//Token is missing, returns with error code 403 Unauthorized
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
		return ""
	}
	return tokenString
}
