package productsmodel

import (
	"context"
	"time"
	"farmsale_backend/config/mdb"

	"go.mongodb.org/mongo-driver/bson"
)

func AllProducts() ([]bson.M, error) {
	var prods = []bson.M{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := mdb.Products.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cur.All(ctx, &prods); err != nil {
		return nil, err
	}

	return prods, nil
}
