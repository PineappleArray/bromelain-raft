package bromelainraft

type raft struct {
	isLeader bool
	myID     int
	timer    time
}

type rpc struct {
}

type server struct {
	id      string
	address string
}

type config struct {
	serverList []server
	total      int
}

type command struct {
}

type log struct {
	commandLog []command
}
