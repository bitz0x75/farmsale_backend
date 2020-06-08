package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/maxwellgithinji/farmsale_backend/controllers/userscontroller"
	"github.com/maxwellgithinji/farmsale_backend/controllers/productscontroller"
)

func main() {
	fmt.Println("Go routines before", runtime.NumCgoCall())
	port := os.Getenv("PORT")
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", index)
	http.HandleFunc("/products", productscontroller.Index)
	http.HandleFunc("/signup", userscontroller.Signup)

	log.Println(http.ListenAndServe(":"+port, nil))
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
