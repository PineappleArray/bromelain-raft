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
	Status           StatusCode   `json:"Status"`
	Payload          PayloadError `json:"Payload"`
	VoteResponse     bool         `json:"VoteResponse"`
	SuccessfulAppend bool         `json:"SuccessfulAppend"`
}

type VoteRequest struct {
	CandidateID string  `json:"CandidateID"`
	Term        int     `json:"Term"`
	LogLen      int     `json:"LogLen"`
	Entry       Command `json:"Entry"`
	RecipientID string  `json:"RecipientID"`
}

type AppendRequest struct {
	entry       Command
	index       int
	term        int
	senderID    string
	recipientID string
}

func (rpc *Rpc) requestVote(request VoteRequest) Response {
	return Response{}
}

func (rpc *Rpc) AppendEntry(request AppendRequest) Response {
	return Response{}
}
