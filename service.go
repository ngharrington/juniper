package main

import (
	"github.com/gorilla/mux"
)

// func handleHome(w http.ResponseWriter, r *http.Request) {
// 	username := r.Context().Value("username")
// 	if username == "" {
// 		username = "anonymous"
// 	}
// 	fmt.Fprintf(w, "Hello, world and %s", username)
// }

type Service struct {
	PrimaryRouter *mux.Router
}

func NewService(router *mux.Router) *Service {
	service := &Service{
		PrimaryRouter: router,
	}
	service.PrimaryRouter.HandleFunc("/", handleHome)
	return service
}
