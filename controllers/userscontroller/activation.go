package userscontroller

import (
	"context"
	"encoding/json"
	"farmsale_backend/config/mdb"
	"farmsale_backend/models/usersmodel"
	"farmsale_backend/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

//ActivateDeactivateAccount is a safer option than deleting accounts which have interracted with the application
//This can be done by admin or owner of the account
func ActivateDeactivateAccount(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()

	user := &usersmodel.User{}

	var users []*usersmodel.User

	params := mux.Vars(req)

	var email = params["email"]

	filter := bson.D{{"email", email}}

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}

	//find the user
	filterCursor, err := mdb.Users.Find(ctx, bson.M{"email": email})
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
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
		return
	}

	if len(users) == 0 {
		err := ErrorResponse{
			Err: `User with email (` + email + `) not found`,
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	//TODO: replace search user by ID so they can update their email
	var update bson.D
	var msg utils.MessageResponse

	if users[0].Isactive {
		update = bson.D{{"$set",
			bson.D{
				{"isactive", false},
			}}}

		msg = utils.MessageResponse{
			Msg: "Account deactivation successful",
		}

		updateUser, err := mdb.Users.UpdateOne(ctx, filter, update)
		if err != nil {
			fmt.Println(err)
			err := ErrorResponse{
				Err: `Update Failed`,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}

		if updateUser.MatchedCount != 0 {
			fmt.Println("matched and replaced an existing document dd") //TODO: delete in prod
			json.NewEncoder(w).Encode(msg)
			return
		}
		if updateUser.UpsertedCount != 0 {
			fmt.Printf("inserted a new document with ID dd %v\n", updateUser.UpsertedID) //TODO: delete in prod
		}
		json.NewEncoder(w).Encode(msg)

	} else {

		update = bson.D{{"$set",
			bson.D{
				{"isactive", true},
			}}}

		msg = utils.MessageResponse{
			Msg: "Account activation successful",
		}

		updateUser, err := mdb.Users.UpdateOne(ctx, filter, update)
		if err != nil {
			fmt.Println(err)
			err := ErrorResponse{
				Err: `Update Failed`,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}

		if updateUser.MatchedCount != 0 {
			fmt.Println("matched and replaced an existing document") //TODO: delete in prod
			json.NewEncoder(w).Encode(msg)
			return
		}
		if updateUser.UpsertedCount != 0 {
			fmt.Printf("inserted a new document with ID %v\n", updateUser.UpsertedID) //TODO: delete in prod
		}
		json.NewEncoder(w).Encode(msg)
	}
}

//DeactivateAccount is a safer option than deleting accounts which have interracted with the application
//This can be done by the owner of the account
//TODO: Remember to log out user in the front end after deactivation
func DeactivateAccount(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()

	user := &usersmodel.User{}

	var users []*usersmodel.User

	params := mux.Vars(req)

	var email = params["email"]

	filter := bson.D{{"email", email}}

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}

	//find the user
	filterCursor, err := mdb.Users.Find(ctx, bson.M{"email": email})
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
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
		return
	}

	if len(users) == 0 {
		err := ErrorResponse{
			Err: `User with email (` + email + `) not found`,
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	//TODO: replace search user by ID so they can update their email
	update := bson.D{{"$set",
		bson.D{
			{"isactive", false},
		}}}

	msg := utils.MessageResponse{
		Msg: "Account deactivation successful",
	}

	updateUser, err := mdb.Users.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: `Update Failed`,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	if updateUser.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document") //TODO: delete in prod
		json.NewEncoder(w).Encode(msg)
		return
	}
	if updateUser.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", updateUser.UpsertedID) //TODO: delete in prod
	}
	json.NewEncoder(w).Encode(msg)
}
