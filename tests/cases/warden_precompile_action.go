package cases

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/warden-protocol/wardenprotocol/precompiles/act"
	"github.com/warden-protocol/wardenprotocol/precompiles/warden"
	"github.com/warden-protocol/wardenprotocol/tests/framework"
	"github.com/warden-protocol/wardenprotocol/tests/framework/exec"
	types "github.com/warden-protocol/wardenprotocol/warden/x/warden/types/v1beta3"
	"testing"
	"time"
)

func init() {
	Register(&Test_WardenPrecompileAction{})
}

type Test_WardenPrecompileAction struct {
	w *exec.WardenNode
}

func (c *Test_WardenPrecompileAction) Setup(t *testing.T, ctx context.Context, build framework.BuildResult) {
	c.w = exec.NewWardenNode(t, build.Wardend)

	go c.w.Start(t, ctx, "./testdata/snapshot-many-users")
	c.w.WaitRunnning(t)
	c.w.PrintDebugLogsAtTheEnd(t, ctx)
}

func (c *Test_WardenPrecompileAction) Run(t *testing.T, ctx context.Context, build framework.BuildResult) {
	alice := exec.NewWardend(c.w, "alice")
	bob := exec.NewWardend(c.w, "bob")
	//dave := exec.NewWardend(c.w, "dave")

	//client := TestGRPCClient(*c.w.GRPCClient(t))
	evmClient := c.w.EthClient(t)

	iActClient, err := act.NewIAct(common.HexToAddress(act.PrecompileAddress), evmClient)
	require.NoError(t, err)

	iWardenClient, err := warden.NewIWarden(common.HexToAddress(warden.PrecompileAddress), evmClient)
	require.NoError(t, err)

	t.Run("work with new action", func(t *testing.T) {
		// addSpaceOwner
		_, err := iWardenClient.AddSpaceOwner(
			alice.TransactOps(t, context.Background(), evmClient),
			1,
			bob.EthAddress(t),
			0,
			0,
			"any(1, warden.space.owners)",
			"any(1, warden.space.owners)")

		require.NoError(t, err)
		time.Sleep(4 * time.Second) // TODO AT: replace by require.Eventually

		actions1, err := iActClient.Actions(alice.CallOps(t), act.TypesPageRequest{})

		require.NoError(t, err)
		require.Len(t, actions1.Actions, 1)

		spaceById1, err2 := iWardenClient.SpaceById(alice.CallOps(t), 1)
		require.NoError(t, err2)
		require.Equal(t, warden.Space{
			Id:                     1,
			Creator:                alice.Address(t),
			Nonce:                  1,
			ApproveSignTemplateId:  0,
			RejectAdminTemplateId:  0,
			ApproveAdminTemplateId: 0,
			RejectSignTemplateId:   0,
			Owners:                 []string{alice.Address(t), bob.Address(t)},
		}, spaceById1)

		// removeSpaceOwner
		_, err = iWardenClient.RemoveSpaceOwner(
			alice.TransactOps(t, context.Background(), evmClient),
			1,
			bob.EthAddress(t),
			1,
			0,
			"any(1, warden.space.owners)",
			"any(1, warden.space.owners)")

		require.NoError(t, err)
		time.Sleep(4 * time.Second) // TODO AT: replace by require.Eventually

		actions2, err := iActClient.Actions(alice.CallOps(t), act.TypesPageRequest{})

		require.NoError(t, err)
		require.Len(t, actions2.Actions, 2)

		spaceById2, err2 := iWardenClient.SpaceById(alice.CallOps(t), 1)
		require.NoError(t, err2)
		require.Equal(t, warden.Space{
			Id:                     1,
			Creator:                alice.Address(t),
			Nonce:                  2,
			ApproveSignTemplateId:  0,
			RejectAdminTemplateId:  0,
			ApproveAdminTemplateId: 0,
			RejectSignTemplateId:   0,
			Owners:                 []string{alice.Address(t)},
		}, spaceById2)

		_, err = iWardenClient.NewKeychain(alice.TransactOps(t, context.Background(), evmClient), "test keychain", warden.KeychainFees{}, "", "", "")

		require.NoError(t, err)
		time.Sleep(4 * time.Second) // TODO AT: replace by require.Eventually

		keychains1, err := iWardenClient.Keychains(alice.CallOps(t), warden.PageRequest{})

		require.NoError(t, err)
		require.Len(t, keychains1.Keychain, 1)

		_, err = iWardenClient.AddKeychainWriter(
			alice.TransactOps(t, context.Background(), evmClient),
			1,
			alice.EthAddress(t))

		require.NoError(t, err)
		time.Sleep(4 * time.Second) // TODO AT: replace by require.Eventually

		keychains2, err := iWardenClient.Keychains(alice.CallOps(t), warden.PageRequest{})

		require.NoError(t, err)
		require.Equal(t, keychains2.Keychain[0].Writers[0], alice.Address(t))

		// newKeyRequest
		_, err = iWardenClient.NewKeyRequest(
			alice.TransactOps(t, context.Background(), evmClient),
			1,
			1,
			uint8(types.KeyType_KEY_TYPE_ECDSA_SECP256K1),
			0,
			0,
			[]warden.Coin{},
			2,
			0,
			"any(1, warden.space.owners)",
			"any(1, warden.space.owners)")

		require.NoError(t, err)
		time.Sleep(4 * time.Second) // TODO AT: replace by require.Eventually

		actions3, err := iActClient.Actions(alice.CallOps(t), act.TypesPageRequest{})

		require.NoError(t, err)
		require.Len(t, actions3.Actions, 3)

		keyRequests1, err := iWardenClient.KeyRequests(alice.CallOps(t), warden.PageRequest{}, 1, int32(types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING), 1)

		require.NoError(t, err)
		require.Len(t, keyRequests1.Keys, 1)
		require.Equal(t, warden.KeyRequest{
			Id:                   1,
			Creator:              alice.Address(t),
			SpaceId:              1,
			KeychainId:           1,
			ApproveTemplateId:    0,
			RejectTemplateId:     0,
			Status:               int32(types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING),
			KeyType:              int32(types.KeyType_KEY_TYPE_ECDSA_SECP256K1),
			DeductedKeychainFees: []warden.Coin{},
			RejectReason:         "",
		}, keyRequests1.Keys[0])

		_, err = iWardenClient.FulfilKeyRequest(
			alice.TransactOps(t, context.Background(), evmClient),
			1,
			[]byte{3, 127, 233, 231, 7, 1, 37, 58, 229, 52, 192, 74, 160, 180, 120, 109, 158, 81, 170, 197, 189, 110, 90, 124, 50, 198, 188, 78, 49, 207, 247, 159, 237},
		)

		require.NoError(t, err)
		time.Sleep(4 * time.Second) // TODO AT: replace by require.Eventually

		keyRequests2, err := iWardenClient.KeyRequests(alice.CallOps(t), warden.PageRequest{}, 1, int32(types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED), 1)

		require.NoError(t, err)
		require.Len(t, keyRequests2.Keys, 1)

		// newSignRequest
		_, err = iWardenClient.NewSignRequest(
			alice.TransactOps(t, context.Background(), evmClient),
			1,
			[]byte{30, 134, 120, 103, 230, 84, 237, 151, 116, 242, 69, 17, 228, 215, 27, 180, 86, 107, 152, 98, 133, 215, 201, 146, 4, 157, 189, 118, 13, 42, 35, 142},
			[]common.Address{},
			[]byte{},
			[]warden.Coin{},
			2,
			0,
			"any(1, warden.space.owners)",
			"any(1, warden.space.owners)")

		require.NoError(t, err)
		time.Sleep(4 * time.Second) // TODO AT: replace by require.Eventually

		actions4, err := iActClient.Actions(alice.CallOps(t), act.TypesPageRequest{})

		require.NoError(t, err)
		require.Len(t, actions4.Actions, 4)

		signRequests, err := iWardenClient.SignRequests(alice.CallOps(t), warden.PageRequest{}, 1, int32(types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING))

		require.NoError(t, err)
		require.Len(t, signRequests.SignRequests, 1)
		require.Equal(t, warden.SignRequest{
			Id:                   1,
			Creator:              alice.Address(t),
			KeyId:                1,
			Status:               int32(types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING),
			EncryptionKey:        []byte{},
			DataForSigning:       []byte{30, 134, 120, 103, 230, 84, 237, 151, 116, 242, 69, 17, 228, 215, 27, 180, 86, 107, 152, 98, 133, 215, 201, 146, 4, 157, 189, 118, 13, 42, 35, 142},
			DeductedKeychainFees: []warden.Coin{},
			Result:               []byte{},
		}, signRequests.SignRequests[0])

		// updateKey
		_, err = iActClient.NewTemplate(
			alice.TransactOps(t, context.Background(), evmClient),
			"test template",
			"any(2, warden.space.owners)")

		require.NoError(t, err)
		time.Sleep(4 * time.Second) // TODO AT: replace by require.Eventually

		_, err = iWardenClient.UpdateKey(
			alice.TransactOps(t, context.Background(), evmClient),
			1,
			1,
			1,
			0,
			"any(1, warden.space.owners)",
			"any(1, warden.space.owners)")

		require.NoError(t, err)
		time.Sleep(4 * time.Second) // TODO AT: replace by require.Eventually

		actions5, err := iActClient.Actions(alice.CallOps(t), act.TypesPageRequest{})

		require.NoError(t, err)
		require.Len(t, actions5.Actions, 5)

		key, err := iWardenClient.KeyById(
			alice.CallOps(t),
			1,
			[]int32{})

		require.NoError(t, err)

		require.Equal(t, warden.Key{
			Id:                1,
			SpaceId:           1,
			ApproveTemplateId: 1,
			RejectTemplateId:  1,
			KeychainId:        1,
			Type:              int32(types.KeyType_KEY_TYPE_ECDSA_SECP256K1),
			PublicKey:         []byte{3, 127, 233, 231, 7, 1, 37, 58, 229, 52, 192, 74, 160, 180, 120, 109, 158, 81, 170, 197, 189, 110, 90, 124, 50, 198, 188, 78, 49, 207, 247, 159, 237},
		}, key.Key)

		// updateSpace
		_, err = iWardenClient.UpdateSpace(
			alice.TransactOps(t, context.Background(), evmClient),
			1,
			2,
			1,
			1,
			1,
			1,
			0,
			"any(1, warden.space.owners)",
			"any(1, warden.space.owners)")

		require.NoError(t, err)
		time.Sleep(4 * time.Second) // TODO AT: replace by require.Eventually

		actions6, err := iActClient.Actions(alice.CallOps(t), act.TypesPageRequest{})

		require.NoError(t, err)
		require.Len(t, actions6.Actions, 6)

		space, err := iWardenClient.SpaceById(
			alice.CallOps(t),
			1)

		require.NoError(t, err)
		require.Equal(t, warden.Space{
			Id:                     1,
			Creator:                alice.Address(t),
			Nonce:                  3,
			ApproveSignTemplateId:  1,
			RejectAdminTemplateId:  1,
			ApproveAdminTemplateId: 1,
			RejectSignTemplateId:   1,
			Owners:                 []string{alice.Address(t)},
		}, space)
	})
}
