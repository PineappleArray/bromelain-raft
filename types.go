package bromelainraft

import (
	"sync"
	"time"
)

// const state that each node can have
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

// contains a list of peers as well as the seed that starts vote
type Config struct {
	serverList []Server
	total      int
	seedVal    int
}

type Command struct {
	data     string
	lampTime int
}

// log that contains the commands
type Log struct {
	commandLog []Command
}
