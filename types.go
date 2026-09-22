package bromelainraft

import "time"

type Raft struct {
	isLeader  bool
	myID      int
	timer     time.Time
	nextIndex map[string]Command
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
