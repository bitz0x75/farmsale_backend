package userscontroller

//TODO: Consider refactoring
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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//EditProfile is only accessible to the owners of the credentials
func EditProfile(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
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

	
	update := bson.D{{"$set",
		bson.D{
			{"username", user.Username},
			{"email", user.Email},
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
