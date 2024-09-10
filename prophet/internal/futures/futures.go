package futures

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type ID uint64

type (
	Input  []byte
	Output []byte
)

type Future struct {
	ID      ID
	Handler string
	Input   Input
}

type FutureResult struct {
	Future
	Output Output
}

type Vote struct {
	ID  ID
	Err error
}

type FutureHandler interface {
	Execute(ctx context.Context, input Input) (Output, error)

	Verify(ctx context.Context, input Input, output Output) error
}

func Run(ctx context.Context, f Future) (FutureResult, error) {
	s := Get(f.Handler)
	if s == nil {
		return FutureResult{}, fmt.Errorf("unknown future: %s", f.Handler)
	}

	log := slog.With("module", "futures", "mode", "run", "future", f.Handler)
	log.Debug("executing", "future", f.ID)
	start := time.Now()
	output, err := s.Execute(ctx, f.Input)
	if err != nil {
		return FutureResult{}, fmt.Errorf("executing future: %w", err)
	}
	log.Debug("done executing", "future", f.ID, "took", time.Since(start))

	return FutureResult{
		Future: f,
		Output: output,
	}, nil
}

func Verify(ctx context.Context, f FutureResult) error {
	s := Get(f.Handler)
	if s == nil {
		return fmt.Errorf("unknown future: %s", f.Handler)
	}

	log := slog.With("module", "futures", "mode", "verify", "future", f.Handler)
	log.Debug("verifying", "proposal", f.ID)
	start := time.Now()
	err := s.Verify(ctx, f.Input, f.Output)
	if err != nil {
		return err
	}
	log.Debug("done verifying", "proposal", f.ID, "took", time.Since(start))

	return nil
}
