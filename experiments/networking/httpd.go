package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Server struct {
	db DB
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.postHandler(w, r)
		return
	case http.MethodGet:
		s.getHandler(w, r)
		return
	}
	http.Error(w, "bad method", http.StatusMethodNotAllowed)
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	data := s.db.Get(key)
	if data == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Write(data)
}

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	defer r.Body.Close()
	rdr := io.LimitedReader{r.Body, 1 << 10}
	data, err := io.ReadAll(&rdr)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}
	s.db.Set(key, data)
	fmt.Fprintf(w, "%s set\n", key)
}

func main() {
	var s Server

	http.HandleFunc("/", s.Handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
