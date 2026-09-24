package bromelainraft

type rpcResponse struct {
	e        error
	response Response
}

type Response struct {
	status int
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

func (rpc *Rpc) requestVote(request VoteRequest) {

}

func (rpc *Rpc) AppendEntry(resquest AppendRequest) {

}
