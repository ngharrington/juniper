package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	PrimaryRouter *mux.Router
	Api           *JournalApi
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}

func (s Service) listJournalEntries(w http.ResponseWriter, r *http.Request) {
	data := s.Api.ListJournalEntries()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (s Service) createJournalEntry(w http.ResponseWriter, r *http.Request) {
	var e JournalEntry
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println(err)
	}
	s.Api.CreateJournalEntry(e.Contents)
}

func NewService(router *mux.Router, api *JournalApi) *Service {
	service := &Service{
		PrimaryRouter: router,
		Api:           api,
	}
	service.PrimaryRouter.HandleFunc("/entry", service.listJournalEntries).Methods("GET")
	service.PrimaryRouter.HandleFunc("/entry", service.createJournalEntry).Methods("POST")
	service.PrimaryRouter.HandleFunc("/", handleHome)
	return service
}
