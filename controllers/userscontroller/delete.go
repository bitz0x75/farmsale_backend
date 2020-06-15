package userscontroller

import (
	"context"
	"encoding/json"
	"github.com/maxwellgithinji/farmsale_backend/config/mdb"
	"github.com/maxwellgithinji/farmsale_backend/models/usersmodel"
	"github.com/maxwellgithinji/farmsale_backend/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteUser is an admin only functionality and this completely removes user accounts
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()

	user := &usersmodel.User{}

	var users []*usersmodel.User

	params := mux.Vars(req)

	//id from params
	strID := params["id"]

	//Convert the id to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(strID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}

	//filter by the id
	filter := bson.D{{"_id", id}}

	err = json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}

	// find the user
	filterCursor, err := mdb.Users.Find(ctx, bson.M{"_id": id})
	if err != nil {
		err := ErrorResponse{
			Err: "ID is invalid",
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
			Err: `User with id (` + strID + `) not found`,
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
