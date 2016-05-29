package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kelseyhightower/memq/broker"
)

var bkr *broker.Broker

type CreateQueueRequest struct {
	Name string `json:"name"`
}

type DeleteQueueRequest struct {
	Name string `json:"name"`
}

type DrainQueueRequest struct {
	Name string `json:"name"`
}

type PutMessageRequest struct {
	Queue string `json:"queue"`
	Body  string `json:"body"`
}

type GetMessageRequest struct {
	Queue string `json:"queue"`
}

func SetBroker(b *broker.Broker) {
	bkr = b
}

func CreateQueueHandler(w http.ResponseWriter, r *http.Request) {
	var cqr CreateQueueRequest
	if err := json.NewDecoder(r.Body).Decode(&cqr); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bkr.CreateQueue(cqr.Name)
}

func DeleteQueueHandler(w http.ResponseWriter, r *http.Request) {
	var dqr DeleteQueueRequest
	if err := json.NewDecoder(r.Body).Decode(&dqr); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bkr.DeleteQueue(dqr.Name)
}

func DrainQueueHandler(w http.ResponseWriter, r *http.Request) {
	var dqr DrainQueueRequest
	if err := json.NewDecoder(r.Body).Decode(&dqr); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bkr.DrainQueue(dqr.Name)
}

func PutMessageHandler(w http.ResponseWriter, r *http.Request) {
	var pmr PutMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&pmr); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m, err := broker.NewMessage(pmr.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := bkr.PutMessage(pmr.Queue, m); err != nil {
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
	var gmr GetMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&gmr); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m, err := bkr.GetMessage(gmr.Queue)
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
