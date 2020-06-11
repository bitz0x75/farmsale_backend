package main

import (
	"farmsale_backend/routes"
	"log"
	"net/http"
	_ "net/http/pprof" // For dev only, dont push to production
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.Handle("/", routes.RouteHandlers())
	log.Println(http.ListenAndServe(":"+port, nil))
}
