package routes

import (
	"github.com/maxwellgithinji/farmsale_backend/middleware/common"
	"net/http"
	_ "net/http/pprof" // For dev only, dont push to production

	"github.com/gorilla/mux"
)

func RouteHandlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	var api = r.PathPrefix("/api").Subrouter()

	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	api.Handle("/favicon.ico", http.NotFoundHandler()).Methods("GET")
	api.Use(common.CommonMiddleware)

	//API V1
	apiV1(api)

	//API V2 can go here
	// apiV2(api)

	return r
}
