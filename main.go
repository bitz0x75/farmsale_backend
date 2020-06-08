package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	"github.com/maxwellgithinji/farmsale_backend/controllers/productscontroller"
	"github.com/maxwellgithinji/farmsale_backend/controllers/userscontroller"
)

func main() {
	fmt.Println("Go routines before", runtime.NumCgoCall())
	port := os.Getenv("PORT")
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", index)
	http.HandleFunc("/products", productscontroller.Index)
	http.HandleFunc("/signup", userscontroller.Signup)

	log.Println(http.ListenAndServe(":"+port, nil))
	fmt.Println("Go routines After", runtime.NumCgoCall())
	// pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	// if err := agent.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	// time.Sleep(time.Hour)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
