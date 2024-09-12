package app

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/warden-protocol/wardenprotocol/prophet/client"
	acttypes "github.com/warden-protocol/wardenprotocol/warden/x/act/types/v1beta1"
	asynctypes "github.com/warden-protocol/wardenprotocol/warden/x/async/types/v1beta1"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	cometabci "github.com/cometbft/cometbft/abci/types"
	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	oraclepreblock "github.com/skip-mev/slinky/abci/preblock/oracle"
	"github.com/skip-mev/slinky/abci/proposals"
	"github.com/skip-mev/slinky/abci/strategies/aggregator"
	"github.com/skip-mev/slinky/abci/strategies/codec"
	compression "github.com/skip-mev/slinky/abci/strategies/codec"
	"github.com/skip-mev/slinky/abci/strategies/currencypair"
	"github.com/skip-mev/slinky/abci/ve"
	vetypes "github.com/skip-mev/slinky/abci/ve/types"
	"github.com/skip-mev/slinky/cmd/constants/marketmaps"
	oracleconfig "github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/pkg/math/voteweighted"
	oracleclient "github.com/skip-mev/slinky/service/clients/oracle"
	servicemetrics "github.com/skip-mev/slinky/service/metrics"
	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
	wardenve "github.com/warden-protocol/wardenprotocol/warden/x/ve/types/v1beta1"
)

func (app *App) initializeOracle(appOpts types.AppOptions) {
	// Read general config from app-opts, and construct oracle service.
	cfg, err := oracleconfig.ReadConfigFromAppOpts(appOpts)
	if err != nil {
		panic(err)
	}

	// If app level instrumentation is enabled, then wrap the oracle service with a metrics client
	// to get metrics on the oracle service (for ABCI++). This will allow the instrumentation to track
	// latency in VerifyVoteExtension requests and more.
	oracleMetrics, err := servicemetrics.NewMetricsFromConfig(cfg, app.ChainID())
	if err != nil {
		panic(err)
	}

	// Create the oracle service.
	app.oracleClient, err = oracleclient.NewClientFromConfig(
		cfg,
		app.Logger().With("client", "oracle"),
		oracleMetrics,
	)
	if err != nil {
		panic(err)
	}

	// Connect to the oracle service (default timeout of 5 seconds).
	go func() {
		if err := app.oracleClient.Start(context.Background()); err != nil {
			app.Logger().Error("failed to start oracle client", "err", err)
			panic(err)
		}

		app.Logger().Info("started oracle client", "address", cfg.OracleAddress)
	}()
	initializeABCIExtensions(app, oracleMetrics)
}

func initializeABCIExtensions(app *App, oracleMetrics servicemetrics.Metrics) {
	veCodec := &AsyncSlinkyVECodec{
		slinkyCodec: compression.NewCompressionVoteExtensionCodec(
			compression.NewDefaultVoteExtensionCodec(),
			compression.NewZLibCompressor(),
		),
	}

	// Create the proposal handler that will be used to fill proposals with
	// transactions and oracle data.
	proposalHandler := proposals.NewProposalHandler(
		app.Logger(),
		baseapp.NoOpPrepareProposal(),
		baseapp.NoOpProcessProposal(),
		ve.NewDefaultValidateVoteExtensionsFn(app.StakingKeeper),
		veCodec,
		compression.NewCompressionExtendedCommitCodec(
			compression.NewDefaultExtendedCommitCodec(),
			compression.NewZStdCompressor(),
		),
		currencypair.NewDeltaCurrencyPairStrategy(app.OracleKeeper),
		oracleMetrics,
	)
	app.SetProcessProposal(
		app.processProposal(
			proposalHandler.ProcessProposalHandler(),
		),
	)

	// Create the aggregation function that will be used to aggregate oracle data
	// from each validator.
	aggregatorFn := voteweighted.MedianFromContext(
		app.Logger(),
		app.StakingKeeper,
		voteweighted.DefaultPowerThreshold,
	)

	// Create the pre-finalize block hook that will be used to apply oracle data
	// to the state before any transactions are executed (in finalize block).
	oraclePreBlockHandler := oraclepreblock.NewOraclePreBlockHandler(
		app.Logger(),
		aggregatorFn,
		app.OracleKeeper,
		oracleMetrics,
		currencypair.NewDeltaCurrencyPairStrategy(app.OracleKeeper),
		veCodec,
		compression.NewCompressionExtendedCommitCodec(
			compression.NewDefaultExtendedCommitCodec(),
			compression.NewZStdCompressor(),
		),
	)

	app.SetPreBlocker(
		app.preBlocker(
			oraclePreBlockHandler.WrappedPreBlocker(app.ModuleManager),
		),
	)

	// Create the vote extensions handler that will be used to extend and verify
	// vote extensions (i.e. oracle data).
	cps := currencypair.NewDeltaCurrencyPairStrategy(app.OracleKeeper)
	extCommitCodec := compression.NewCompressionExtendedCommitCodec(
		compression.NewDefaultExtendedCommitCodec(),
		compression.NewZStdCompressor(),
	)
	voteExtensionsHandler := ve.NewVoteExtensionHandler(
		app.Logger(),
		app.oracleClient,
		time.Second,
		cps,
		veCodec,
		aggregator.NewOraclePriceApplier(
			aggregator.NewDefaultVoteAggregator(
				app.Logger(),
				aggregatorFn,
				// we need a separate price strategy here, so that we can optimistically apply the latest prices
				// and extend our vote based on these prices
				currencypair.NewDeltaCurrencyPairStrategy(app.OracleKeeper),
			),
			app.OracleKeeper,
			veCodec,
			extCommitCodec,
			app.Logger(),
		),
		oracleMetrics,
	)

	veManager := NewVoteExtensionManager()
	veManager.Register(
		voteExtensionsHandler.ExtendVoteHandler(),
		voteExtensionsHandler.VerifyVoteExtensionHandler(),
		proposalHandler.PrepareProposalHandler(),
	)
	veManager.Register(
		func(ctx sdk.Context, req *cometabci.RequestExtendVote) (*cometabci.ResponseExtendVote, error) {
			// TODO: add my votes fetched from prophet
			// use AsyncVoteExtension proto
			votes, err := client.FetchVotes(context.TODO())
			if err != nil {
				ctx.Logger().Error("failed to fetch votes", "err", err)
				return &cometabci.ResponseExtendVote{
					VoteExtension: []byte{},
				}, nil
			}

			asyncVotes := make([]*asynctypes.AsyncVoteExtensionItem, len(votes.Result))
			for i, v := range votes.Result {
				status := asynctypes.FutureVoteType_VOTE_TYPE_REJECTED
				if v.Approved {
					status = asynctypes.FutureVoteType_VOTE_TYPE_VERIFIED
				}
				asyncVotes[i] = &asynctypes.AsyncVoteExtensionItem{
					FutureId: v.ID,
					Vote:     status,
				}
			}

			asyncve := asynctypes.AsyncVoteExtension{
				Votes: asyncVotes,
			}
			asyncveBytes, err := asyncve.Marshal()
			if err != nil {
				return nil, err
			}

			return &cometabci.ResponseExtendVote{
				VoteExtension: asyncveBytes,
			}, nil
		},
		func(ctx sdk.Context, req *cometabci.RequestVerifyVoteExtension) (*cometabci.ResponseVerifyVoteExtension, error) {
			// TODO: check req.VoteExtension is a valid AsyncVoteExtension
			// (eg. is it a valid proto)
			return &cometabci.ResponseVerifyVoteExtension{
				Status: cometabci.ResponseVerifyVoteExtension_ACCEPT,
			}, nil
		},
		app.xasyncPrepareProposal(),
	)

	app.SetPrepareProposal(veManager.PrepareProposalHandler())
	app.SetExtendVoteHandler(veManager.ExtendVoteHandler())
	app.SetVerifyVoteExtensionHandler(veManager.VerifyVoteExtensionHandler())
}

type AppUpgrade struct {
	Name         string
	Handler      upgradetypes.UpgradeHandler
	StoreUpgrade storetypes.StoreUpgrades
}

// createSlinkyUpgrader returns the upgrade name and an upgrade handler that:
// - runs migrations
// - updates the consensus keeper params with a vote extension enable height. (height of upgrade + 10).
// - adds the core markets to x/marketmap keeper.
// additionally, it returns the required StoreUpgrades needed for the new slinky modules added to this chain.
func createSlinkyUpgrader(app *App) AppUpgrade {
	return AppUpgrade{
		Name: "v03-to-v04",
		Handler: func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			// renamed module
			fromVM[acttypes.ModuleName] = fromVM["intent"]
			delete(fromVM, "intent")

			migrations, err := app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
			if err != nil {
				return nil, err
			}
			// upgrade consensus params to enable vote extensions
			consensusParams, err := app.ConsensusParamsKeeper.Params(ctx, nil)
			if err != nil {
				return nil, err
			}
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			consensusParams.Params.Abci = &tmtypes.ABCIParams{
				VoteExtensionsEnableHeight: sdkCtx.BlockHeight() + int64(10),
			}
			_, err = app.ConsensusParamsKeeper.UpdateParams(ctx, &consensustypes.MsgUpdateParams{
				Authority: app.ConsensusParamsKeeper.GetAuthority(),
				Block:     consensusParams.Params.Block,
				Evidence:  consensusParams.Params.Evidence,
				Validator: consensusParams.Params.Validator,
				Abci:      consensusParams.Params.Abci,
			})
			if err != nil {
				return nil, err
			}

			// add core markets
			coreMarkets := marketmaps.CoreMarketMap
			markets := coreMarkets.Markets
			keys := make([]string, 0, len(markets))
			for name := range markets {
				keys = append(keys, name)
			}
			slices.Sort(keys)

			// iterates over slice and not map
			for _, marketName := range keys {
				// create market
				market := markets[marketName]
				err = app.MarketMapKeeper.CreateMarket(sdkCtx, market)
				if err != nil {
					return nil, err
				}

				// invoke hooks
				err = app.MarketMapKeeper.Hooks().AfterMarketCreated(sdkCtx, market)
				if err != nil {
					return nil, err
				}
			}

			err = app.MarketMapKeeper.SetParams(
				sdkCtx,
				marketmaptypes.Params{
					Admin: authtypes.NewModuleAddress(govtypes.ModuleName).String(), // governance. allows gov to add or remove market authorities.
					// market authority addresses may add and update markets to the x/marketmap module.
					MarketAuthorities: []string{
						"warden1ua63s43u2p4v38pxhcxmps0tj2gudyw2rctzk3",          // skip multisig
						authtypes.NewModuleAddress(govtypes.ModuleName).String(), // governance
					}},
			)
			if err != nil {
				return nil, fmt.Errorf("failed to set x/marketmap params: %w", err)
			}

			return migrations, nil
		},
		StoreUpgrade: storetypes.StoreUpgrades{
			Added: []string{
				marketmaptypes.ModuleName,
				oracletypes.ModuleName,
			},
		},
	}
}

func (app *App) xasyncPrepareProposal() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *cometabci.RequestPrepareProposal) (*cometabci.ResponsePrepareProposal, error) {
		resp := &cometabci.ResponsePrepareProposal{
			Txs: req.Txs,
		}

		if !ve.VoteExtensionsEnabled(ctx) {
			return resp, nil
		}

		log := ctx.Logger().With("module", "prophet")
		asyncTx, err := buildAsyncTx(ctx, req.LocalLastCommit.Votes)
		if err != nil {
			log.Error("failed to build async tx", "err", err)
			return resp, nil
		}
		resp.Txs = trimExcessBytes(resp.Txs, req.MaxTxBytes-int64(len(asyncTx)))
		ctx.Logger().Info("injecting x/async tx", "num txs before", len(resp.Txs))
		resp.Txs = injectTx(asyncTx, 1, resp.Txs)

		return resp, nil
	}
}

func (app *App) processProposal(process sdk.ProcessProposalHandler) sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *cometabci.RequestProcessProposal) (*cometabci.ResponseProcessProposal, error) {
		resp, err := process(ctx, req)
		if err != nil || resp.Status == cometabci.ResponseProcessProposal_REJECT {
			return resp, err
		}

		if !ve.VoteExtensionsEnabled(ctx) || len(req.Txs) < 2 {
			return resp, nil
		}

		log := ctx.Logger().With("module", "prophet")
		asyncTx := req.Txs[1]
		if len(asyncTx) == 0 {
			return resp, nil
		}

		var tx asynctypes.AsyncInjectedTx
		if err := tx.Unmarshal(asyncTx); err != nil {
			log.Error("failed to unmarshal async tx", "err", err)
			// probably not an async tx?
			// but slinky in this case rejects their proposal so maybe we
			// should do the same?
			return &cometabci.ResponseProcessProposal{
				Status: cometabci.ResponseProcessProposal_ACCEPT,
			}, nil
		}

		return &cometabci.ResponseProcessProposal{
			Status: cometabci.ResponseProcessProposal_ACCEPT,
		}, nil
	}
}

func (app *App) preBlocker(preBlock sdk.PreBlocker) sdk.PreBlocker {
	return func(ctx sdk.Context, req *cometabci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
		resp, err := preBlock(ctx, req)
		if err != nil {
			return resp, err
		}

		if !ve.VoteExtensionsEnabled(ctx) || len(req.Txs) < 2 {
			return resp, nil
		}

		log := ctx.Logger().With("module", "prophet")
		asyncTx := req.Txs[1]
		if len(asyncTx) == 0 {
			return resp, nil
		}

		var tx asynctypes.AsyncInjectedTx
		if err := tx.Unmarshal(asyncTx); err != nil {
			log.Error("failed to unmarshal async tx", "err", err)
			// probably not an async tx?
			// but slinky in this case rejects their proposal so maybe we
			// should do the same?
			return resp, nil
		}

		proposer := ctx.BlockHeader().ProposerAddress
		for _, r := range tx.Results {
			if err := app.AsyncKeeper.AddFutureResult(ctx, r.Id, proposer, r.Output); err != nil {
				return resp, err
			}
		}

		for _, v := range tx.ExtendedVotesInfo {
			var w wardenve.VoteExtensions
			if err := w.Unmarshal(v.VoteExtension); err != nil {
				return resp, fmt.Errorf("failed to unmarshal vote extension wrapper: %w", err)
			}
			// todo: check VE signature, or maybe do it in the verify ve handler?
			if len(w.Extensions) < 2 {
				continue
			}
			var asyncve asynctypes.AsyncVoteExtension
			if err := asyncve.Unmarshal(w.Extensions[1]); err != nil {
				return resp, fmt.Errorf("failed to unmarshal x/async vote extension: %w", err)
			}
			for _, vote := range asyncve.Votes {
				// v.Validator.Address is the Comet validator address
				// i.e. the one with `wardenvalcons` prefix
				if err := app.AsyncKeeper.SetFutureVote(ctx, vote.FutureId, v.Validator.Address, vote.Vote); err != nil {
					return resp, fmt.Errorf("failed to set task vote: %w", err)
				}
			}
		}

		return resp, nil
	}
}

func buildAsyncTx(ctx context.Context, votes []cometabci.ExtendedVoteInfo) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()
	res, err := client.FetchCompletedFutures(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*asynctypes.FutureResult, len(res.Result))
	for i, r := range res.Result {
		results[i] = &asynctypes.FutureResult{
			Id:     r.ID,
			Output: r.Output,
		}
	}

	tx := asynctypes.AsyncInjectedTx{
		Results:           results,
		ExtendedVotesInfo: votes,
	}

	txBytes, err := tx.Marshal()
	if err != nil {
		return nil, err
	}

	return txBytes, nil
}

func injectTx(newTx []byte, position int, appTxs [][]byte) [][]byte {
	if position < 0 {
		panic("position must be >= 0")
	}

	if position == 0 {
		return append([][]byte{newTx}, appTxs...)
	}

	if position >= len(appTxs) {
		return append(appTxs, newTx)
	}

	return append(appTxs[:position], append([][]byte{newTx}, appTxs[position:]...)...)
}

func trimExcessBytes(txs [][]byte, maxSizeBytes int64) [][]byte {
	var (
		returnedTxs   [][]byte
		consumedBytes int64
	)
	for _, tx := range txs {
		consumedBytes += int64(len(tx))
		if consumedBytes > maxSizeBytes {
			break
		}
		returnedTxs = append(returnedTxs, tx)
	}
	return returnedTxs
}

func lenBytes(txs [][]byte) int {
	var l int
	for _, tx := range txs {
		l += len(tx)
	}
	return l
}

///

type AsyncSlinkyVECodec struct {
	slinkyCodec codec.VoteExtensionCodec
}

var _ compression.VoteExtensionCodec = (*AsyncSlinkyVECodec)(nil)

func (a *AsyncSlinkyVECodec) Decode(b []byte) (vetypes.OracleVoteExtension, error) {
	var w wardenve.VoteExtensions
	if err := w.Unmarshal(b); err != nil {
		return vetypes.OracleVoteExtension{}, err
	}

	return a.slinkyCodec.Decode(w.Extensions[0])
}

func (a *AsyncSlinkyVECodec) Encode(ve vetypes.OracleVoteExtension) ([]byte, error) {
	return a.slinkyCodec.Encode(ve)
}
