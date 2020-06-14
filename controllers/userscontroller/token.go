package userscontroller

import (
	"encoding/json"
	"farmsale_backend/models/jwtmodel"
	"farmsale_backend/models/usersmodel"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateToken(w http.ResponseWriter, user *usersmodel.User) {
	now := time.Now()

	tk := &jwtmodel.Token{
		Username:      user.Username,
		Email:         user.Email,
		Phonenumber:   user.Phonenumber,
		Idnumber:      user.Idnumber,
		Userclass:     user.Userclass,
		Isadmin:       user.Isadmin,
		Isvalid:       user.Isvalid,
		Isblacklisted: user.Isblacklisted,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: now.Add(time.Minute * 100000).Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Add(-5 * time.Second).Unix(),
			Issuer:    os.Getenv("BASE_URL"),
			Subject:   "AccessToken",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	// decodestr, _ := base64.URLEncoding.DecodeString(os.Getenv("TOKEN_SECRET"))
	decodebs := []byte(os.Getenv("TOKEN_SECRET"))
	tokenString, err := token.SignedString(decodebs)
	if err != nil {
		err := ErrorResponse{
			Err: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
		return
	}

	//Store user token
	var resp = map[string]interface{}{}
	resp["token"] = tokenString
	resp["User"] = tk
	json.NewEncoder(w).Encode(resp)
}
