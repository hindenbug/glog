package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Log *Log
}

type WriteRequest struct {
	Record Record `json:"record"`
}

type WriteResponse struct {
	Offset uint64 `json:"offset"`
}

type ReadRequest struct {
	Offset uint64 `json:"offset"`
}

type ReadResponse struct {
	Record Record `json:"record"`
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
	var req WriteRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset, err := s.Log.Append(req.Record)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := WriteResponse{Offset: offset}
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *httpServer) handleRead(w http.ResponseWriter, r *http.Request) {
	var req ReadRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, err := s.Log.Read(req.Offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := ReadResponse{Record: record}
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
