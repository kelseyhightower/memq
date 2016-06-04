package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kelseyhightower/memq/broker"
)

var bkr *broker.Broker

func SetBroker(b *broker.Broker) {
	bkr = b
}

func CreateQueueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bkr.CreateQueue(vars["name"])
}

func DeleteQueueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bkr.DeleteQueue(vars["name"])
}

func DrainQueueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bkr.DrainQueue(vars["name"])
}

func PutMessageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	m, err := broker.NewMessage(string(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := bkr.PutMessage(vars["name"], m); err != nil {
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

func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m, err := bkr.GetMessage(vars["name"])
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

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	s := bkr.Stats()
	data, err := json.MarshalIndent(&s, "", "  ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
