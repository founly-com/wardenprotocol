package http

import (
	"log/slog"
	"net/http"

	"github.com/warden-protocol/wardenprotocol/prophet/internal/futures"
)

type FutureResultStorage interface {
	Take(n int) ([]futures.FutureResult, error)
	Ack(ids []futures.ID) error
}

type FutureResultStorageDebug interface {
	PendingFutures() ([]futures.FutureResult, error)
}

type VoteStorage interface {
	Take(n int) ([]futures.Vote, error)
	Ack(ids []futures.ID) error
}

type VoteStorageDebug interface {
	PendingVotes() ([]futures.Vote, error)
}

type Server struct {
	log       *slog.Logger
	addr      string
	sink      FutureResultStorage
	votesSink VoteStorage
}

func NewServer(addr string, sink FutureResultStorage, votesSink VoteStorage) *Server {
	return &Server{
		log:       slog.With("module", "http"),
		addr:      addr,
		sink:      sink,
		votesSink: votesSink,
	}
}

func (s *Server) Serve() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /future-results", s.getFutureResults)
	mux.HandleFunc("POST /ack-futures", s.postAck)
	mux.HandleFunc("GET /votes", s.getVotes)

	debugMux := http.NewServeMux()
	debugMux.HandleFunc("GET /debug/pending-future-results", s.getPendingFutures)
	debugMux.HandleFunc("GET /debug/pending-votes", s.getPendingVotes)
	mux.Handle("/debug/", debugMux)

	return http.ListenAndServe(s.addr, mux)
}
