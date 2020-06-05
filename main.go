package main

import (
	"farmsale_backend/controllers/userscontroller"
	"net/http"
	"os"

	"github.com/maxwellgithinji/farmsale_backend/controllers/productscontroller"
)

func main() {
	port := os.Getenv("PORT")
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// http.HandleFunc("/", index)
	http.HandleFunc("/products", productscontroller.Index)
	http.HandleFunc("/signup", userscontroller.Signup)

	http.ListenAndServe(":"+port, nil)
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "/products", http.StatusSeeOther)
// }
