package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Log *Log
}

func newHTTPServer() *httpServer {
	return &httpServer{Log: NewLog()}
}

func NewHttpServer(addr string) *http.Server {
	server := newHTTPServer()
	r := mux.NewRouter()

	r.HandleFunc("/", server.handleWrite).Methods("POST")
	r.HandleFunc("/", server.handleRead).Methods("GET")

	return &http.Server{Addr: addr, Handler: r}
}

func (s *httpServer) handleWrite(w http.ResponseWriter, r *http.Request) {

}

func (s *httpServer) handleRead(w http.ResponseWriter, r *http.Request) {

}
