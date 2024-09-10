package wardenai

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/warden-protocol/wardenprotocol/prophet/internal/futures"
)

func init() {
	futures.Register("wardenai", Future{})
}

type Future struct {
}

func (s Future) Execute(ctx context.Context, input futures.Input) (futures.Output, error) {
	res, err := http.Post("http://localhost:9001/job/solve", "application/json", bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s Future) Verify(ctx context.Context, input futures.Input, output futures.Output) error {
	// todo: verify output
	return nil
}
