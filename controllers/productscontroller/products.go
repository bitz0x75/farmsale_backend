package productscontroller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/maxwellgithinji/farmsale_backend/config/mdb"
	"github.com/maxwellgithinji/farmsale_backend/models/productsmodel"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

var DB = mdb.ConnectDB()
var Products = DB.Collection("products")

func Index(w http.ResponseWriter, req *http.Request) {

	var prods = &productsmodel.Product{}

	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := Products.Find(ctx, prods)
	if err != nil {
		err := ErrorResponse{
			Err: "Error finding products",
		}
		json.NewEncoder(w).Encode(err)
		return
	}
	if err = cur.All(ctx, &prods); err != nil {
		err := ErrorResponse{
			Err: "Error finding products",
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prods)
}
