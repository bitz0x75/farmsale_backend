package userscontroller

import (
	"context"
	"encoding/json"
	"farmsale_backend/config/mdb"
	"farmsale_backend/models/usersmodel"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()

	user := &usersmodel.User{}

	var users []*usersmodel.User

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	//validate user login details
	if string(user.Password) == "" {
		err := ErrorResponse{
			Err: "Please enter password",
		}
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err)
		return
	}
	if user.Email == "" {
		err := ErrorResponse{
			Err: "Please enter email",
		}
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err)
		return
	}

	//find the user
	filterCursor, err := mdb.Users.Find(ctx, bson.M{"email": user.Email})
	if err != nil {
		err := ErrorResponse{
			Err: "Email is invalid",
		}
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
		return
	}

	if err = filterCursor.All(ctx, &users); err != nil {
		err := ErrorResponse{
			Err: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
		return
	}

	if len(users) == 0 {
		err := ErrorResponse{
			Err: `User with email (` + user.Email + `) not found`,
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}

	//decode pssword
	pass := users[0].Password

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		err := ErrorResponse{
			Err: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
		return
	}

	//Add access token on login
	generateToken(w, users[0])
}
