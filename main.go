// Copyright 2016 Google, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kelseyhightower/memq/api"
	"github.com/kelseyhightower/memq/broker"
)

func main() {
	api.SetBroker(broker.New())
	r := mux.NewRouter()
	r.HandleFunc("/stats", api.StatsHandler).Methods("GET")
	r.HandleFunc("/queues/{name}", api.CreateQueueHandler).Methods("POST")
	r.HandleFunc("/queues/{name}", api.DeleteQueueHandler).Methods("DELETE")
	r.HandleFunc("/queues/{name}/drain", api.DrainQueueHandler).Methods("POST")
	r.HandleFunc("/queues/{name}/messages", api.GetMessageHandler).Methods("GET")
	r.HandleFunc("/queues/{name}/messages", api.PutMessageHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
