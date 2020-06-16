package routes

import (
	"encoding/json"
	"github.com/maxwellgithinji/farmsale_backend/utils"
	"net/http"
)

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
