package client

import (
	"context"

	types "github.com/warden-protocol/wardenprotocol/warden/x/async/types/v1beta1"
	"google.golang.org/grpc"
)

type AsyncQueryClient struct {
	client types.QueryClient
}

// NewWardenQueryClient returns a WardenQueryClient
func NewAsyncQueryClient(c *grpc.ClientConn) *AsyncQueryClient {
	return &AsyncQueryClient{
		client: types.NewQueryClient(c),
	}
}

func (t *AsyncQueryClient) PendingFutures(ctx context.Context, page *PageRequest) ([]types.Future, error) {
	res, err := t.client.PendingFutures(ctx, &types.QueryPendingFuturesRequest{
		Pagination: page,
	})
	if err != nil {
		return nil, err
	}

	return res.Futures, nil
}

func (t *AsyncQueryClient) FuturesPendingVote(ctx context.Context, validator string, page *PageRequest) ([]types.FutureWithResultResponse, error) {
	res, err := t.client.FuturesPendingVote(ctx, &types.QueryFuturesPendingVoteRequest{
		Validator:  validator,
		Pagination: page,
	})
	if err != nil {
		return nil, err
	}

	return res.Futures, nil
}
