package userscontroller

import (
	"context"
	"encoding/json"
	"farmsale_backend/models/usersmodel"
	"fmt"
	"net/http"
	"regexp"

	"github.com/maxwellgithinji/farmsale_backend/config/mdb"
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

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Invalid data provided",
		}
		fmt.Println(req.Body, string(user.Password))
		json.NewEncoder(w).Encode(err)
		return
	}

	// // validate request body values
	// if string(user.Password) == "" {
	// 	err := ErrorResponse{
	// 		Err: "Please enter password",
	// 	}
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }
	if user.Username == "" || user.Email == "" || user.Phonenumber == "" || user.Idnumber == 0 {
		err := ErrorResponse{
			Err: "All fields must be complete",
		}
		fmt.Println(req.Body, user)

		json.NewEncoder(w).Encode(err)
		return
	}

	//Se model indexes
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

		json.NewEncoder(w).Encode(err)
		return
	}

	//encrypt the password
	bs, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		err := ErrorResponse{
			Err: "Password encryption failed",
		}

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
		//11000 is mongodb error code for duplicate content
		//Note we have created unique indexes for username and email
		if errCode == 11000 {
			err := ErrorResponse{
				Err: "Username or email already exists",
			}

			json.NewEncoder(w).Encode(err)
			return
		}
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	json.NewEncoder(w).Encode(cur)
}
