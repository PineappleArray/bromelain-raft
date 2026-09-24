package bromelainraft

type rpcResponse struct {
	e        error
	response Response
}

type Response struct {
}

func (rpc *Rpc) requestVote(candidateID string, term int) {

}
