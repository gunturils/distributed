package raft

import "sync"

type Role string

const (
	Leader    Role = "leader"
	Follower  Role = "follower"
	Candidate Role = "candidate"
)

type Node struct {
	mu       sync.Mutex
	ID       string
	Role     Role
	Term     int
	VotedFor string   // which node ID this node voted for in the current term
	Peers    []string // addresses of other nodes, e.g. "http://localhost:8081"
}

func NewNode(id string, peers []string) *Node {
	return &Node{
		ID:       id,
		Role:     Follower, // everyone starts as a follower now
		Term:     0,
		VotedFor: "",
		Peers:    peers,
	}
}
