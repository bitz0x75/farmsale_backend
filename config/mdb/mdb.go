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
)

var (
	//DB is our db initialization
	DB *mongo.Database
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() *mongo.Database {
	connstr := os.Getenv("DB_CONN_STR")
	connstr=`mongodb+srv://maxgit:icuifourcu4@farmsale-hh9n7.gcp.mongodb.net/test?retryWrites=true&w=majority`
	client, err := mongo.NewClient(options.Client().ApplyURI(connstr))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	// defer client.Disconnect(ctx)

	//Select db according to the current environment
	os.Getenv("ENV")
	if os.Getenv("ENV") == "development" {
		DB = client.Database("farmsaleDev")
		fmt.Println("You connected to DEV database") //For demo only
		return DB
	} else if os.Getenv("ENV") == "production" {
		DB = client.Database("farmsale")
		fmt.Println("You connected to PROD database") //For demo only
		return DB
	} else if os.Getenv("ENV") == "testing" {
		DB = client.Database("farmsaleTest")
		fmt.Println("You connected to TEST database") //For demo only
		return DB
	} else {
		err := os.Setenv("ENV", "development")
		if err != nil {
			log.Fatal("Error setting env variable")
			return nil
		}
		DB = client.Database("farmsaleDev")
		fmt.Println("You connected to DEV database") //For demo only
		return DB
	}
}

var db = ConnectDB()
var Users = db.Collection("users")
var Products = DB.Collection("products")
