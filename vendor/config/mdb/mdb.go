package mdb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	//DB is our db initialization
	DB *mongo.Database

	//Products is a collection in the db
	Products *mongo.Collection
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	loadDB()
}

func loadDB() {
	//initialize client
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_CONN_STR")))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	// defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())

	//Select db according to the current environment
	os.Getenv("ENV")
	if os.Getenv("ENV") == "development" {
		DB = client.Database("farmsaleDev")
	} else if os.Getenv("ENV") == "production" {
		DB = client.Database("farmsale")
	} else if os.Getenv("ENV") == "testing" {
		DB = client.Database("farmsaleTest")
	} else {
		err := os.Setenv("ENV", "development")
		if err != nil {
			log.Fatal("Error setting env variable")
		}
		DB = client.Database("farmsaleDev")
	}

	Products = DB.Collection("products")

	fmt.Println("You connected to your database")
}
