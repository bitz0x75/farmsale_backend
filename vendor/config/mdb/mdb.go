package mdb

import (
	"github.com/globalsign/mgo"
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)


var (
	//DB is our db initialization
	// DB  *mongo.Database
	DB *mgo.Database

	//Products is a collection in the db
	// Products *mongo.Collection
	Products *mgo.Collection


)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConn := os.Getenv("DEV_DB_CONN_STR")
	s, err := mgo.Dial(dbConn)
	if err != nil {
		fmt.Println("db connection failed")
		panic(err)
	}
	
	DB = s.DB("farmsale")

	Products = DB.C("products")
	
	fmt.Println("You connected to your database")
}

