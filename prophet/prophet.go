package prophet

import (
	"github.com/warden-protocol/wardenprotocol/prophet/internal/exec"
	"github.com/warden-protocol/wardenprotocol/prophet/internal/ingress"
	"github.com/warden-protocol/wardenprotocol/prophet/types"
)

func RunFutureLoop(src ingress.FutureSource, sink types.FutureResultWriter) {
	if err := exec.Futures(src, sink); err != nil {
		panic(err)
	}
}

func RunFutureVotesLoop(proposalSrc ingress.FutureResultSource, sink types.VoteWriter) {
	if err := exec.Votes(proposalSrc, sink); err != nil {
		panic(err)
	}
}
