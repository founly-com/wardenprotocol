package exec

import (
	"context"
	"log/slog"

	"github.com/warden-protocol/wardenprotocol/prophet/internal/futures"
	"github.com/warden-protocol/wardenprotocol/prophet/internal/ingress"
)

type FutureResultWriter interface {
	Add(result futures.FutureResult) error
}

func Futures(src ingress.FutureSource, sink FutureResultWriter) error {
	log := slog.With("module", "pipe_future_request")

	reqs, err := ingress.Futures(src)
	if err != nil {
		return err
	}
	futs := mapchan(reqs, mapFuture)

	go func() {
		for future := range futs {
			tlog := log.With("future", future.ID)

			tlog.Debug("running future")
			output, err := futures.Run(context.TODO(), future)
			if err != nil {
				tlog.Error("failed to run future", "err", err)
				continue
			}
			err = sink.Add(output)
			if err != nil {
				tlog.Error("failed to add future to sink", "err", err)
				continue
			}
		}
	}()

	return nil
}

type VoteWriter interface {
	Add(result futures.Vote) error
}

func Votes(src ingress.FutureResultSource, sink VoteWriter) error {
	log := slog.With("module", "pipe_verify_proposal")

	reqs, err := ingress.FutureResults(src)
	if err != nil {
		return err
	}
	proposals := mapchan(reqs, mapFutureResult)

	go func() {
		for proposal := range proposals {
			plog := log.With("proposal", proposal.ID)

			plog.Debug("verifying proposal")
			err := futures.Verify(context.TODO(), proposal)
			if err := sink.Add(futures.Vote{
				ID:  proposal.ID,
				Err: err,
			}); err != nil {
				plog.Error("failed to add future to sink", "err", err)
				continue
			}
		}
	}()

	return nil
}

func mapchan[A, B any](reqs <-chan A, fn func(A) B) <-chan B {
	futures := make(chan B)
	go func() {
		defer close(futures)
		for req := range reqs {
			futures <- fn(req)
		}
	}()
	return futures
}

func mapFuture(req ingress.Future) futures.Future {
	return futures.Future{
		ID:      futures.ID(req.ID),
		Handler: req.Handler,
		Input:   req.Input,
	}
}

func mapFutureResult(req ingress.FutureResult) futures.FutureResult {
	return futures.FutureResult{
		Future: mapFuture(req.Future),
		Output: req.Output,
	}
}
