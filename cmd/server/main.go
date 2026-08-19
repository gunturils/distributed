package main

import (
	"log"
	"net/http"

	"github.com/gunturils/distributed/internal/api"
	"github.com/gunturils/distributed/internal/raft"
	"github.com/gunturils/distributed/internal/store"
)

func main() {
	s := &api.Server{
		Store: store.New(),
		Node:  raft.NewNode("node-1"),
	}

	http.HandleFunc("/status", s.StatusHandler)
	http.HandleFunc("/key", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.GetKeyHandler(w, r)
		} else {
			s.SetKeyHandler(w, r)
		}
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
