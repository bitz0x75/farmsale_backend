package productscontroller

import (
	"encoding/json"
	"fmt"
	"models/productsmodel"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	prods, err := productsmodel.AllProducts()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	prodj, err := json.Marshal(prods)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", prodj)

}
