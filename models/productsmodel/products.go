package productsmodel

type Product struct {
	// ID       bson.ObjectId // `json:"id" bson:"_id"`
	Item         string // `json:"item" bson:"item"`
	Currentstock float64 // `json:"currentstock" bson:"email"`
	standard     string // `json:"standard" bson:"standard"`
	Unitprice    float64 // `json:"unitprice" bson:"unitprice"`
}
