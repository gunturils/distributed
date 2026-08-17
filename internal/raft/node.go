package raft

type Role string

const (
	Leader   Role = "leader"
	Follower Role = "follower"
)

type Node struct {
	ID   string
	Role Role
	Term int
}

func NewNode(id string) *Node {
	return &Node{ID: id, Role: Leader, Term: 0} // hardcoded for now
}
