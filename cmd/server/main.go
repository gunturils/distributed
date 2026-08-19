package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/gunturils/distributed/internal/api"
	"github.com/gunturils/distributed/internal/raft"
	"github.com/gunturils/distributed/internal/store"
)

func main() {
	id := flag.String("id", "node-1", "unique ID for this node")
	port := flag.String("port", "8080", "port this node listens on")
	peers := flag.String("peers", "", "comma-separated list of peer addresses, e.g. http://localhost:8081,http://localhost:8082")
	flag.Parse()

	var peerList []string
	if *peers != "" {
		peerList = strings.Split(*peers, ",")
	}

	s := &api.Server{
		Store: store.New(),
		Node:  raft.NewNode(*id, peerList),
	}

	http.HandleFunc("/status", s.StatusHandler)
	http.HandleFunc("/key", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.GetKeyHandler(w, r)
		} else {
			s.SetKeyHandler(w, r)
		}
	})

	log.Printf("node %s listening on :%s (peers: %v)", *id, *port, peerList)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
