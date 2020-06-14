package jwtmodel

import jwt "github.com/dgrijalva/jwt-go"

type Token struct {
	Username      string // `json:"username" bson:"username"`
	Email         string // `json:"email" bson:"email"`
	Phonenumber   string // `json:"phonenumber" bson:"phonenumber"`
	Idnumber      int    // `json:"idnumber" bson:"idnumber"`
	Userclass     string // `json:"userclass" bson:"userclass"`
	Isadmin       bool   // `json:"isadmin" bson:"isadmin"`
	Isblacklisted bool   // `json:"isblacklisted" bson:"isblacklisted"`
	Isvalid       bool   // `json:"isvalid" bson:"isvalid"`
	*jwt.StandardClaims
}
