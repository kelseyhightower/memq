// Copyright 2016 Google, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kelseyhightower/memq/api"
	"github.com/kelseyhightower/memq/broker"
)

var (
	listenAddr string
)

func main() {
	flag.StringVar(&listenAddr, "http", "0.0.0.0:80", "HTTP listen address.")
	flag.Parse()

	log.Printf("memq server starting...")
	log.Printf("listening on %s", listenAddr)

	s := api.NewServer(broker.New())

	r := mux.NewRouter()
	r.HandleFunc("/stats", s.StatsHandler).Methods("GET")
	r.HandleFunc("/queues/{name}", s.CreateQueueHandler).Methods("POST")
	r.HandleFunc("/queues/{name}", s.DeleteQueueHandler).Methods("DELETE")
	r.HandleFunc("/queues/{name}/drain", s.DrainQueueHandler).Methods("POST")
	r.HandleFunc("/queues/{name}/messages", s.GetMessageHandler).Methods("GET")
	r.HandleFunc("/queues/{name}/messages", s.PutMessageHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(listenAddr, r))
}
