package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Service *Service
	Router  *mux.Router
}

type RequestLogger struct {
	handler http.Handler
}

func (l *RequestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	proxyIpAddr := r.RemoteAddr
	forwardedForIpAddre := r.Header.Get("x-forwarded-for")
	log.Printf("%s %s %v %s %s", r.Method, r.URL.Path, time.Since(start), proxyIpAddr, forwardedForIpAddre)
}

func NewRequestLogger(handlerToWrap http.Handler) *RequestLogger {
	return &RequestLogger{handlerToWrap}
}

func NewServer() *Server {
	router := mux.NewRouter()
	store := GetJournalEntryStore("memory")
	api := &JournalApi{
		store: store,
	}
	service := NewService(router, api)
	return &Server{
		Service: service,
		Router:  router,
	}
}

func (s *Server) StartServer() {
	addr := fmt.Sprintf("localhost:8080")
	log.Printf("server running at %s", addr)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
