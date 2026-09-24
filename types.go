package bromelainraft

import (
	"sync"
	"time"
)

type lState string

const (
	leader    lState = "leader"
	follower  lState = "follower"
	candidate lState = "candidate"
)

type Raft struct {
	mu          sync.Mutex
	myID        string
	timer       time.Time
	leaderState lState

	// persistent state
	currentTerm int
	votedFor    string
	log         Log

	nextIndex  map[string]int
	matchIndex map[string]int

	config Config
}

type Rpc struct {
	sender   string
	reciever string
}

type Vote struct {
	voterID     string
	candidateID string
	time        time.Time
}

type Server struct {
	id      string
	address string
}

type Config struct {
	serverList []Server
	total      int
	seedVal    int
}

type Command struct {
	data     string
	lampTime int
}

type Log struct {
	commandLog []Command
}
