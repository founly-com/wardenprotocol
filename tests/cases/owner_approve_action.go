package cases

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/warden-protocol/wardenprotocol/warden/x/act/types/v1beta1"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/warden-protocol/wardenprotocol/tests/framework"
	"github.com/warden-protocol/wardenprotocol/tests/framework/checks"
	"github.com/warden-protocol/wardenprotocol/tests/framework/exec"
	types "github.com/warden-protocol/wardenprotocol/warden/x/warden/types/v1beta3"
)

func init() {
	Register(&Test_OwnerApproveAction{})
}

type Test_OwnerApproveAction struct {
	w *exec.WardenNode
}

func (c *Test_OwnerApproveAction) Setup(t *testing.T, ctx context.Context, build framework.BuildResult) {
	c.w = exec.NewWardenNode(t, build.Wardend)

	go c.w.Start(t, ctx, "./testdata/snapshot-many-users")
	c.w.WaitRunnning(t)
}

func (c *Test_OwnerApproveAction) Run(t *testing.T, ctx context.Context, _ framework.BuildResult) {
	client := TestGRPCClient(*c.w.GRPCClient(t))

	alice := exec.NewWardend(c.w, "alice")
	bob := exec.NewWardend(c.w, "bob")
	charlie := exec.NewWardend(c.w, "charlie")
	dave := exec.NewWardend(c.w, "dave")

	addNewOwnerCommandTemplate := "warden new-action add-space-owner --space-id %d --new-owner %s"

	resAddOwner := bob.Tx(t, fmt.Sprintf(addNewOwnerCommandTemplate, 1, bob.Address(t)))
	checks.SuccessTx(t, resAddOwner)
	client.EnsureSpaceAmount(t, ctx, bob.Address(t), 0)

	resApproveBob := alice.Tx(t, "act approve-action --action-id 1")
	checks.SuccessTx(t, resApproveBob)
	client.EnsureSpaceAmount(t, ctx, bob.Address(t), 1)

	resNewRule := alice.Tx(t, "act new-rule --name approve_requires_two --definition \"any(2, warden.space.owners)\"")
	checks.SuccessTx(t, resNewRule)

	resUpdateSpaceAdminRuleByCharlie := charlie.Tx(t, "warden new-action update-space --space-id 1 --admin-rule-id 1")
	checks.SuccessTx(t, resUpdateSpaceAdminRuleByCharlie)

	spaceAfterInvalidApprove, err := client.Warden.SpaceById(ctx, &types.QuerySpaceByIdRequest{
		Id: 1,
	})
	require.NoError(t, err)
	require.Equal(t, uint64(0), spaceAfterInvalidApprove.Space.AdminRuleId)

	resUpdateSpaceAdminRuleByAlice := alice.Tx(t, "warden new-action update-space --space-id 1 --admin-rule-id 1")
	checks.SuccessTx(t, resUpdateSpaceAdminRuleByAlice)

	spaceAfterValidApprove, err := client.Warden.SpaceById(ctx, &types.QuerySpaceByIdRequest{
		Id: 1,
	})
	require.NoError(t, err)
	require.Equal(t, uint64(1), spaceAfterValidApprove.Space.AdminRuleId)

	resAliceAddOwnerCharlie := alice.Tx(t, fmt.Sprintf(addNewOwnerCommandTemplate, 1, charlie.Address(t)))
	checks.SuccessTx(t, resAliceAddOwnerCharlie)
	client.EnsureSpaceAmount(t, ctx, charlie.Address(t), 0)

	resApproveCharlie := bob.Tx(t, "act approve-action --action-id 4")
	checks.SuccessTx(t, resApproveCharlie)
	client.EnsureSpaceAmount(t, ctx, charlie.Address(t), 1)

	resCharlieAddOwnerDave := charlie.Tx(t, fmt.Sprintf(addNewOwnerCommandTemplate, 1, dave.Address(t)))
	checks.SuccessTx(t, resCharlieAddOwnerDave)
	client.EnsureSpaceAmount(t, ctx, dave.Address(t), 0)

	resApproveDaveByBob := bob.Tx(t, "act approve-action --action-id 5")
	checks.SuccessTx(t, resApproveDaveByBob)
	client.EnsureSpaceAmount(t, ctx, dave.Address(t), 1)
}

type TestGRPCClient exec.GRPCClient

func (client *TestGRPCClient) EnsureSpaceAmount(t *testing.T, ctx context.Context, ownerKey string, amount int) {
	resSpacesByDaveAfterBobApprove, err := client.Warden.SpacesByOwner(ctx, &types.QuerySpacesByOwnerRequest{
		Owner: ownerKey,
	})
	require.NoError(t, err)
	require.Equal(t, amount, len(resSpacesByDaveAfterBobApprove.Spaces))
}

func (client *TestGRPCClient) EnsureActionStatus(t *testing.T, ctx context.Context, actionId uint64, expectedStatus v1beta1.ActionStatus) {
	res, err := client.Act.ActionById(ctx, &v1beta1.QueryActionByIdRequest{Id: actionId})
	require.NoError(t, err)
	require.Equal(t, expectedStatus, res.Action.Status)
}

func (client *TestGRPCClient) EnsureBalanceAmount(t *testing.T, ctx context.Context, ownerKey string, balance sdk.Coin) {
	balanceResponse, err := client.Bank.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: ownerKey,
		Denom:   balance.Denom,
	})
	require.NoError(t, err)
	require.True(t, balance.Equal(balanceResponse.Balance))
}

func (client *TestGRPCClient) GetBalanceAmount(t *testing.T, ctx context.Context, ownerKey string, denom string) sdk.Coin {
	balanceResponse, err := client.Bank.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: ownerKey,
		Denom:   denom,
	})
	require.NoError(t, err)
	require.NotNil(t, balanceResponse.Balance)

	return *balanceResponse.Balance
}
