package jwtmodel

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID            primitive.ObjectID // `json:"id" bson:"_id"`
	Username      string             // `json:"username" bson:"username"`
	Email         string             // `json:"email" bson:"email"`
	Phonenumber   string             // `json:"phonenumber" bson:"phonenumber"`
	Idnumber      int                // `json:"idnumber" bson:"idnumber"`
	Userclass     string             // `json:"userclass" bson:"userclass"`
	Isadmin       bool               // `json:"isadmin" bson:"isadmin"`
	Isblacklisted bool               // `json:"isblacklisted" bson:"isblacklisted"`
	Isvalid       bool               // `json:"isvalid" bson:"isvalid"`
	Isactive      bool               // `json:"isactive" bson:"isactive"`
	*jwt.StandardClaims
}
