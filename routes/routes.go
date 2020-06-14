package routes

import (
	"encoding/json"
	"farmsale_backend/controllers/productscontroller"
	"farmsale_backend/controllers/userscontroller"
	"farmsale_backend/middleware/auth"
	"farmsale_backend/utils"
	"net/http"
	_ "net/http/pprof" // For dev only, dont push to production

	"github.com/gorilla/mux"
)

func RouteHandlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.Handle("/favicon.ico", http.NotFoundHandler()).Methods("GET")

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/signup", userscontroller.Signup).Methods("POST")
	r.HandleFunc("/login", userscontroller.Login).Methods("POST")

	//Auth Route
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	s.HandleFunc("/products", productscontroller.Index).Methods("GET")
	s.HandleFunc("/profile/{email}", userscontroller.EditProfile).Methods("PUT")

	//Admin Route
	a := r.PathPrefix("/admin").Subrouter()
	a.Use(auth.AdminVerify)
	a.HandleFunc("/", admin).Methods("GET")
	a.HandleFunc("/profile/delete/{email}", userscontroller.DeleteUser).Methods("DELETE")

	//Manager Route
	m := r.PathPrefix("/manager").Subrouter()
	m.Use(auth.ManagerVerify)
	m.HandleFunc("/", manager).Methods("GET")

	//Agent Route
	ag := r.PathPrefix("/agent").Subrouter()
	ag.Use(auth.AgentVerify)
	ag.HandleFunc("/", agent).Methods("GET")
	ag.HandleFunc("/profile/{email}", userscontroller.EditProfile).Methods("PUT")

	//Current user route
	cu := r.PathPrefix("/profile").Subrouter()
	cu.Use(auth.CurrentUserVerify)
	cu.HandleFunc("/{email}", userscontroller.EditProfile).Methods("PUT")

	return r
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	msg := utils.MessageResponse{
		Msg: "User welcome to farmsale",
	}
	json.NewEncoder(w).Encode(msg)
}

func admin(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	msg := utils.MessageResponse{
		Msg: "Admin welcome to farmsale",
	}
	json.NewEncoder(w).Encode(msg)
}

func agent(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	msg := utils.MessageResponse{
		Msg: "Agent welcome to farmsale",
	}
	json.NewEncoder(w).Encode(msg)
}

func manager(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	msg := utils.MessageResponse{
		Msg: "Manager welcome to farmsale",
	}
	json.NewEncoder(w).Encode(msg)
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
