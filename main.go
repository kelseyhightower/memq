// Copyright 2016 Google, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/memq/api"
	"github.com/kelseyhightower/memq/broker"
)

func main() {
	api.SetBroker(broker.New())

	http.HandleFunc("/broker/stats", api.StatsHandler)
	http.HandleFunc("/queue/create", api.CreateQueueHandler)
	http.HandleFunc("/queue/delete", api.DeleteQueueHandler)
	http.HandleFunc("/queue/drain", api.DrainQueueHandler)
	http.HandleFunc("/message/get", api.GetMessageHandler)
	http.HandleFunc("/message/put", api.PutMessageHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
