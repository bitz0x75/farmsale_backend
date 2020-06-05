package userscontroller

import (
	"farmsale_backend/config/mdb"
	"farmsale_backend/models/usersmodel"
	"encoding/json"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"context"
	"time"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

var DB = mdb.ConnectDB()
var Users = DB.Collection("users")

func Signup(w http.ResponseWriter, req *http.Request){
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	user := &usersmodel.User{}

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// validate request body values
	if user.Username == "" || user.Email == "" || user.Password == "" || user.Phonenumber == "" || user.IDnumber == "" {
		err := ErrorResponse{
			Err: "All fields must be complete",
		}
		fmt.Println(req.Body)
		json.NewEncoder(w).Encode(err)
		return
	}

	//encrypt the password
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		err := ErrorResponse{
			Err: "Password encryption failed",
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	//check for valid idnumber datatype


	
	user.Password = string(pass)

	//Insert User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := Users.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cur)
}