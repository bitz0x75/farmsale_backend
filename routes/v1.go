package routes

import (
	"github.com/maxwellgithinji/farmsale_backend/controllers/productscontroller"
	"github.com/maxwellgithinji/farmsale_backend/controllers/userscontroller"
	"github.com/maxwellgithinji/farmsale_backend/middleware/auth"

	"github.com/gorilla/mux"
)

func apiV1(api *mux.Router) {
	var api1 = api.PathPrefix("/v1").Subrouter()

	api1.HandleFunc("/", index).Methods("GET")
	api1.HandleFunc("/signup", userscontroller.Signup).Methods("POST")
	api1.HandleFunc("/login", userscontroller.Login).Methods("POST")

	//Auth Route
	s := api1.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	s.HandleFunc("/products", productscontroller.Index).Methods("GET")
	s.HandleFunc("/profile/{id}", userscontroller.EditProfile).Methods("PUT")

	//Admin Route
	a := api1.PathPrefix("/admin").Subrouter()
	a.Use(auth.AdminVerify)
	a.HandleFunc("/", admin).Methods("GET")
	a.HandleFunc("/profile/delete/{id}", userscontroller.DeleteUser).Methods("DELETE")
	a.HandleFunc("/profile/activate/{id}", userscontroller.ActivateDeactivateAccount).Methods("PUT")

	//Manager Route
	m := api1.PathPrefix("/manager").Subrouter()
	m.Use(auth.ManagerVerify)
	m.HandleFunc("/", manager).Methods("GET")

	//Agent Route
	ag := api1.PathPrefix("/agent").Subrouter()
	ag.Use(auth.AgentVerify)
	ag.HandleFunc("/", agent).Methods("GET")

	//Current user route
	cu := api1.PathPrefix("/profile").Subrouter()
	cu.Use(auth.CurrentUserVerify)
	cu.HandleFunc("/{id}", userscontroller.EditProfile).Methods("PUT")
	cu.HandleFunc("/deactivate/{id}", userscontroller.DeactivateAccount).Methods("PUT")

}
