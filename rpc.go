package bromelainraft

const voteEndPoint = "/raft/request-vote"
const AppendEndPoint = "/raft/append-entries"

type StatusCode string

const (
	badrequest StatusCode = "400 Bad Request"
	notfound   StatusCode = "404 Not Found"
	invalid    StatusCode = "405"
	unexpected StatusCode = "500 Internal Server Error"
	notready   StatusCode = "503 Service Unavailable"
)

type PayloadError string

const (
	ok PayloadError = "200 OK"
)

type Response struct {
	status           StatusCode
	playload         PayloadError
	voteGranted      bool
	successfulAppend bool
}

type VoteRequest struct {
	candidateID string
	term        int
	logLen      int
	entry       Command
	recipientID string
}

type AppendRequest struct {
	entry       Command
	index       int
	term        int
	senderID    string
	recipientID string
}

func (rpc *Rpc) requestVote(request VoteRequest) Response {
	return nil
}

func (rpc *Rpc) AppendEntry(request AppendRequest) Response {
	return nil
}
