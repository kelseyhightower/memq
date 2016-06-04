package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kelseyhightower/memq/broker"

	"github.com/gorilla/mux"
)

type Server struct {
	broker *broker.Broker
}

func NewServer(b *broker.Broker) *Server {
	return &Server{broker: b}
}

func (s *Server) CreateQueueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.broker.CreateQueue(vars["name"])
}

func (s *Server) DeleteQueueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.broker.DeleteQueue(vars["name"])
}

func (s *Server) DrainQueueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.broker.DrainQueue(vars["name"])
}

func (s *Server) PutMessageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	if err := s.broker.PutMessage(vars["name"], string(body)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m, err := s.broker.GetMessage(vars["name"])
	if err == broker.ErrEmptyQueue {
		log.Println(err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.MarshalIndent(&m, "", "  ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (s *Server) StatsHandler(w http.ResponseWriter, r *http.Request) {
	stats := s.broker.Stats()
	data, err := json.MarshalIndent(&stats, "", "  ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
