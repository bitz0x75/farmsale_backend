package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // For dev only, dont push to production
	"os"

	"github.com/maxwellgithinji/farmsale_backend/routes"
)

func main() {
	port := os.Getenv("PORT")
	http.Handle("/", routes.RouteHandlers())
	log.Println(http.ListenAndServe(":"+port, nil))
}
