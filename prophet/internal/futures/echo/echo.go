// Package echo is a mock future that echoes back the input.
package echo

import (
	"bytes"
	"context"
	"fmt"

	"github.com/warden-protocol/wardenprotocol/prophet/internal/futures"
)

func init() {
	futures.Register("echo", Future{})
}

type Future struct{}

func (s Future) Execute(ctx context.Context, input futures.Input) (futures.Output, error) {
	return futures.Output(input), nil
}

func (s Future) Verify(ctx context.Context, input futures.Input, output futures.Output) error {
	if bytes.Compare(input, output) != 0 {
		return fmt.Errorf("input and output do not match")
	}
	return nil
}
