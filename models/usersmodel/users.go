package usersmodel

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type User struct {
	// ID       bson.ObjectId // `json:"id" bson:"_id"`
	Username      string // `json:"username" bson:"username"`
	Email         string // `json:"email" bson:"email"`
	Password      string // `json:"password" bson:"password"`
	Phonenumber   string // `json:"phonenumber" bson:"phonenumber"`
	Idnumber      int    // `json:"idnumber" bson:"idnumber"`
	Userclass     string // `json:"userclass" bson:"userclass"`
	Isadmin       bool   // `json:"isadmin" bson:"isadmin"`
	Isblacklisted bool   // `json:"isblacklisted" bson:"isblacklisted"`
	Isvalid       bool   // `json:"isvalid" bson:"isvalid"`

}

func SetEmailIndex(coll *mongo.Collection) {
	ctx := context.Background()
	indexOptions := options.Index().SetUnique(true)
	indexKeys := bsonx.MDoc{
		"email": bsonx.Int32(-1),
	}
	noteIndexModel := mongo.IndexModel{
		Options: indexOptions,
		Keys:    indexKeys,
	}
	indexName, err := coll.Indexes().CreateOne(ctx, noteIndexModel)
	if err != nil {
		fmt.Println(err, "errrrrr====")
		return
	}
	fmt.Println(indexName) // Output: email_-1
	return
}

func SetUsernameIndex(coll *mongo.Collection) {
	ctx := context.Background()
	indexOptions := options.Index().SetUnique(true)
	indexKeys := bsonx.MDoc{
		"username": bsonx.Int32(-1),
	}
	noteIndexModel := mongo.IndexModel{
		Options: indexOptions,
		Keys:    indexKeys,
	}
	indexName, err := coll.Indexes().CreateOne(ctx, noteIndexModel)
	if err != nil {
		fmt.Println(err, "errrrrr====")
		return
	}
	fmt.Println(indexName) // Output: username_-1
	return
}
