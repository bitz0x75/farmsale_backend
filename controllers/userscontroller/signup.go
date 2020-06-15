package userscontroller

import (
	"context"
	"encoding/json"
	"farmsale_backend/config/mdb"
	"farmsale_backend/models/usersmodel"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

func Signup(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()

	user := &usersmodel.User{}

	//default userclass is user and admin cannot signup normally
	user.Userclass = "user"
	user.Isadmin = false
	user.Isvalid = true        //TODO: implement user verification by email, Unverified users should be deleted after 1 day and cannot be able to interract with other endpoints
	user.Isblacklisted = false //TODO: blacklisted users are those who violate the app policy, they can't interract with the app, but their data is kept since they have interracted with the app
	user.Isactive = true //TODO: a user who is inactive can no longer login but they have already interracted with the app
	
	// validate decoded values
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Invalid data provided",
		}
		fmt.Println(req.Body, string(user.Password))
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err)
		return
	}

	// validate request body values
	if user.Username == "" || string(user.Password) == "" || user.Email == "" || user.Phonenumber == "" || user.Userclass == "" || user.Idnumber == 0 {
		err := ErrorResponse{
			Err: "All fields must be complete",
		}
		fmt.Println(req.Body, user)
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err)
		return
	}

	//model indexes for unique entries
	usersmodel.SetEmailIndex(mdb.Users)
	usersmodel.SetUsernameIndex(mdb.Users)

	//validate password length
	if len(user.Password) < 8 {
		err := ErrorResponse{
			Err: "Password should be at least 8 characters",
		}

		json.NewEncoder(w).Encode(err)
		return
	}

	//Validate email
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(user.Email) {
		err := ErrorResponse{
			Err: "Email is invalid",
		}
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err)
		return
	}

	//encrypt the password
	bs, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		err := ErrorResponse{
			Err: "Password encryption failed",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password = string(bs)

	//Insert User
	cur, err := mdb.Users.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err, "my errr")
		//Initialize exception to be thrown my mongo db duplicate data
		var merr mongo.WriteException
		merr = err.(mongo.WriteException)
		errCode := merr.WriteErrors[0].Code
		//11000 is mongodb error cow.WriteHeader(httw.WriteHeader(http.StatusForbidden)p.StatusForbidden)de for duplicate content
		//Note we have created unique indexes for username and email
		if errCode == 11000 {
			err := ErrorResponse{
				Err: "Username or email already exists",
			}
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(err)
			return
		}
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}
	log.Println(cur)
	//Generate token on signup
	generateToken(w, user)
}
