package auth

import (
	"encoding/json"
	"farmsale_backend/utils"
	"net/http"
	"os"
	"strings"
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

