package productscontroller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/maxwellgithinji/farmsale_backend/models/productsmodel"
	"github.com/maxwellgithinji/farmsale_backend/config/mdb"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

func Index(w http.ResponseWriter, req *http.Request) {

	var prods = &productsmodel.Product{}

	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	ctx := context.Background()
	cur, err := mdb.Products.Find(ctx, prods)
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
