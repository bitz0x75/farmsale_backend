package userscontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/maxwellgithinji/farmsale_backend/config/mdb"
	"github.com/maxwellgithinji/farmsale_backend/models/usersmodel"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

var DB = mdb.ConnectDB()
var Users = DB.Collection("users")

// func init() {
// 	// TODO: create a go routine for tasks taking long
	
// }

func Signup(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := &usersmodel.User{}

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Invalid data provided",
		}
		fmt.Println(req.Body)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	// validate request body values
	if user.Username == "" || user.Email == "" || user.Password == "" || user.Phonenumber == "" || user.Idnumber == 0 {
		err := ErrorResponse{
			Err: "All fields must be complete",
		}
		fmt.Println(req.Body)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	//Validate Pass not nil
	usersmodel.SetEmailIndex(Users)
	usersmodel.SetUsernameIndex(Users)
	//validate password length
	if len(user.Password) < 8 {
		err := ErrorResponse{
			Err: "Password should be at least 8 characters",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	//Validate email
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(user.Email) {
		err := ErrorResponse{
			Err: "Email is invalid",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	//encrypt the password
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		err := ErrorResponse{
			Err: "Password encryption failed",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password = string(pass)

	//Insert User
	cur, err := Users.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err, "my errr")
		//Initialize exception to be thrown my mongo db duplicate data
		var merr mongo.WriteException
		merr = err.(mongo.WriteException)
		errCode := merr.WriteErrors[0].Code
		//11000 is mongodb error code for duplicate content
		//Note we have created unique indexes for username and email
		if errCode == 11000 {
			err := ErrorResponse{
				Err: "Username or email already exists",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cur)
}
