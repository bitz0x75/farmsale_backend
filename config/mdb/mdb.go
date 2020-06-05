package mdb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//DB is our db initialization
	DB *mongo.Database
)

func ConnectDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_CONN_STR")))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	//Select db according to the current environment
	os.Getenv("ENV")
	if os.Getenv("ENV") == "development" {
		DB = client.Database("farmsaleDev")
		fmt.Println("You connected to DEV database")//For demo only
		return DB
	} else if os.Getenv("ENV") == "production" {
		DB = client.Database("farmsale")
		fmt.Println("You connected to PROD database")//For demo only
		return DB
	} else if os.Getenv("ENV") == "testing" {
		DB = client.Database("farmsaleTest")
		fmt.Println("You connected to TEST database")//For demo only
		return DB
	} else {
		err := os.Setenv("ENV", "development")
		if err != nil {
			log.Fatal("Error setting env variable")
			return nil
		}
		DB = client.Database("farmsaleDev")
		fmt.Println("You connected to DEV database")//For demo only
		return DB
	}
}