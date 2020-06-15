package productsmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id           primitive.ObjectID  `json:"id" bson:"_id"`
	Item         string        // `json:"item" bson:"item"`
	Currentstock float64       // `json:"currentstock" bson:"currentstock"`
	Standard     string        // `json:"standard" bson:"standard"`
	Unitprice    float64       // `json:"unitprice" bson:"unitprice"`
}
