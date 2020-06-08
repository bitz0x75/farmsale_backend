package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // For dev only, dont push to production
	"os"

	"github.com/maxwellgithinji/farmsale_backend/controllers/productscontroller"
	"github.com/maxwellgithinji/farmsale_backend/controllers/userscontroller"
)

func main() {
	port := os.Getenv("PORT")
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", index)
	http.HandleFunc("/products", productscontroller.Index)
	http.HandleFunc("/signup", userscontroller.Signup)

	log.Println(http.ListenAndServe(":"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
