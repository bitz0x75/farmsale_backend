package usersmodel


type User struct {
	// ID       bson.ObjectId // `json:"id" bson:"_id"`
	Username    string // `json:"username" bson:"username"`
	Email       string // `json:"email" bson:"email"`
	Password    string // `json:"password" bson:"password"` 
	Phonenumber string // `json:"phonenumber" bson:"phonenumber"`
	IDnumber    string    // `json:"idnumber" bson:"idnumber"`

}
