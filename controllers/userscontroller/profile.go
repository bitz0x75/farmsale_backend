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
	"golang.org/x/crypto/bcrypt"
)

func EditProfile(w http.ResponseWriter, req *http.Request) {
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

	// find the user
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

	//TODO: replace search user by ID so they can update their email
	update := bson.D{{"$set",
		bson.D{
			{"username", user.Username},
			{"password", user.Password},
			{"phonenumber", user.Phonenumber},
			{"idnumber", user.Idnumber},
		}}}

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

	msg := utils.MessageResponse{
		Msg: "Update successful",
	}

	if updateUser.MatchedCount != 0 {
		fmt.Println("matched and replaced " + fmt.Sprint(len(users)) + " existing document") //TODO: delete in prod
		json.NewEncoder(w).Encode(msg)
		return
	}
	if updateUser.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", updateUser.UpsertedID) //TODO: delete in prod
	}
	json.NewEncoder(w).Encode(msg)

}

func BlacklistUser(w http.ResponseWriter, req *http.Request) {
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
			{"isblacklisted", true},
		}}}
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

	msg := utils.MessageResponse{
		Msg: "Update successful",
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

func ActivateAccount(w http.ResponseWriter, req *http.Request) {
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
			{"isactive", true},
		}}}
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

	msg := utils.MessageResponse{
		Msg: "Update successful",
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

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
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
	//TODO: replace search user by ID so they can delete by ID
	deleteUser, err := mdb.Users.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: `Deletion Failed`,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	msg := utils.MessageResponse{
		Msg: "Deletion Successful successful",
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteUser.DeletedCount)
	json.NewEncoder(w).Encode(msg)
}
