package main

import (
	"net/http"
	"os"
	"farmsale_backend/controllers/productscontroller"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", index)
	http.HandleFunc("/products", productscontroller.Index)
	http.ListenAndServe(":"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
