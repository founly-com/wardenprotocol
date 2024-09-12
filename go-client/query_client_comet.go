package client

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc"
)

// WardenQueryClient is the client for the treasury module.
type CometQueryClient struct {
	cc *grpc.ClientConn
}

// NewWardenQueryClient returns a WardenQueryClient
func NewCometQueryClient(c *grpc.ClientConn) *CometQueryClient {
	return &CometQueryClient{
		cc: c,
	}
}

func (q *CometQueryClient) LatestValidatorSet(ctx context.Context) (*cmtservice.GetLatestValidatorSetResponse, error) {
	out := new(cmtservice.GetLatestValidatorSetResponse)
	in := &cmtservice.GetLatestValidatorSetRequest{
		Pagination: &query.PageRequest{
			Limit: 500,
		},
	}
	err := q.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetLatestValidatorSet", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
