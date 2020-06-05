package productsmodel

type Product struct {
	// ID       bson.ObjectId // `json:"id" bson:"_id"`
	Item    string // `json:"item" bson:"item"`
	Currentstock       string // `json:"currentstock" bson:"email"`
	standard    string // `json:"standard" bson:"standard"` 
	Unitprice string // `json:"unitprice" bson:"unitprice"`
}
