package productscontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/maxwellgithinji/farmsale_backend/config/mdb"
	"github.com/maxwellgithinji/farmsale_backend/models/productsmodel"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

func Index(w http.ResponseWriter, req *http.Request) {

	var prods = []productsmodel.Product{}
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	ctx := context.Background()
	cur, err := mdb.Products.Find(ctx, &bson.M{})
	if err != nil {
		fmt.Printf("%+v\n", err)
		err := ErrorResponse{
			Err: "Error finding products",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	err = cur.All(ctx, &prods)
	if err != nil {
		fmt.Printf("%+v\n", err)
		err := ErrorResponse{
			Err: "Error finding all products",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prods)
}
