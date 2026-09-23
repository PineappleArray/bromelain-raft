package bromelainraft

import "time"

type lState string

const (
	leader    lState = "leader"
	follower  lState = "follower"
	candidate lState = "candidate"
)

type Raft struct {
	leaderState lState
	myID        int
	timer       time.Time
	nextIndex   map[string]Command
	currentTerm int
	votedFor    string
	commitIndex int
	lastApplied int
}

type Rpc struct {
}

type Vote struct {
	id   string
	time time.Time
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
