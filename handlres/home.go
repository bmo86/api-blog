package handlres

import (
	"encoding/json"
	"net/http"
	"rest-websockets/server"
)

//home response
type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) //code 200
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to server with go! ",
			Status:  true,
		})
	}
}
