package routes

import (
	"encoding/json"
	"farmsale_backend/config/utils"
	"farmsale_backend/controllers/userscontroller"
	"net/http"
	_ "net/http/pprof" // For dev only, dont push to production

	"github.com/gorilla/mux"
	"github.com/maxwellgithinji/farmsale_backend/controllers/productscontroller"
)

func RouteHandlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.Handle("/favicon.ico", http.NotFoundHandler()).Methods("GET")

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/signup", userscontroller.Signup).Methods("POST")
	r.HandleFunc("/login", userscontroller.Login).Methods("POST")
	r.HandleFunc("/products", productscontroller.Index).Methods("GET")
	
	return r
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	msg := utils.MessageResponse{
		Msg: "Welcome to farmsale",
	}
	json.NewEncoder(w).Encode(msg)
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
