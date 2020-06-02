package main

import (
	"controllers/productscontroller"
	"net/http"
	"os"
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
