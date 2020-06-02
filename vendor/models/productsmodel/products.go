package productsmodel

import (
	"config/mdb"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	// ID     bson.ObjectId // `json:"id" bson:"_id"`
	Item   string  // `json:"item" bson:"item"`
	Currentstock  float32  // `json:"currentstock" bson:"currentstock"`
	Standard string  // `json:"standard" bson:"standard"`
	Unitprice  float32 // `json:"unitprice" bson:"unitprice"`
}


func AllProducts()([]Product, error) {
	prods := []Product{}
	err := mdb.Products.Find(bson.M{}).All(&prods)
	if err != nil {
		return nil, err
	}
	
	return prods, nil
}
