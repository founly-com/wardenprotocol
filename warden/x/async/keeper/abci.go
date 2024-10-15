package keeper

import (
	"context"
)

func (k Keeper) EndBlocker(ctx context.Context) error {
	k.FuturesSource.Flush(ctx, k)
	k.ResultsSource.Flush(ctx, k)
	return nil
}
