package cli

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/warden-protocol/wardenprotocol/warden/app/inference"
	actcli "github.com/warden-protocol/wardenprotocol/warden/x/act/client"
	"github.com/warden-protocol/wardenprotocol/warden/x/warden/types/v1beta3"
)

// NewTxCmd returns a root CLI command handler for x/warden transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        v1beta3.ModuleName,
		Short:                      "Warden transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewActionTxCmd(),
		FulfillKeyRequestTxCmd(),
		RejectKeyRequestTxCmd(),
		FulfillSignRequestTxCmd(),
		RejectSignRequestTxCmd(),
		NewInferenceRequestTxCmd(),
	)

	return txCmd
}

func NewActionTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-action",
		Short: "Create a new Action subcommands",
	}

	cmd.AddCommand(
		actcli.RegisterActionCmd(&v1beta3.MsgAddSpaceOwner{}, "Add a new owner to a Space"),
		actcli.RegisterActionCmd(&v1beta3.MsgNewKeyRequest{}, "Request a new Key"),
		actcli.RegisterActionCmd(&v1beta3.MsgNewSignRequest{}, "Request a signature"),
		actcli.RegisterActionCmd(&v1beta3.MsgRemoveSpaceOwner{}, "Remove an owner from a Space"),
		actcli.RegisterActionCmd(&v1beta3.MsgUpdateKey{}, "Update a Key information"),
		actcli.RegisterActionCmd(&v1beta3.MsgUpdateSpace{}, "Update a Space information"),
	)

	return cmd
}

func FulfillKeyRequestTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fulfill-key-request [request-id] [public-key-data]",
		Short: "Fulfill a key request providing the public key.",
		Long: `Fulfill a key request providing the public key.
The sender of this transaction must be a writer of the Keychain for the request.
The public key must be a base64 encoded string.`,
		Example: fmt.Sprintf("%s tx warden fulfill-key-request 1234 aGV5dGhlcmU=", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			reqId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			pk, err := base64.StdEncoding.DecodeString(args[1])
			if err != nil {
				return err
			}

			msg := &v1beta3.MsgFulfilKeyRequest{
				Creator:   clientCtx.GetFromAddress().String(),
				Status:    v1beta3.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED,
				RequestId: reqId,
				Result:    v1beta3.NewMsgFulfilKeyRequestKey(pk),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func RejectKeyRequestTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reject-key-request [request-id] [reason]",
		Short: "Reject a key request providing the reason.",
		Long: `Reject a key request providing a reason.
The sender of this transaction must be a writer of the Keychain for the request.`,
		Example: fmt.Sprintf("%s tx warden reject-key-request 1234 'something happened'", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			reqId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := &v1beta3.MsgFulfilKeyRequest{
				Creator:   clientCtx.GetFromAddress().String(),
				Status:    v1beta3.KeyRequestStatus_KEY_REQUEST_STATUS_REJECTED,
				RequestId: reqId,
				Result:    v1beta3.NewMsgFulfilKeyRequestReject(args[1]),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func FulfillSignRequestTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fulfill-sign-request [request-id] [sign-data]",
		Short: "Fulfill a signature request providing the signature.",
		Long: `Fulfill a signature request providing the signature.
The sender of this transaction must be a writer of the Keychain for the request.
The sign-data must be a base64 encoded string.`,
		Example: fmt.Sprintf("%s tx warden fulfill-sign-request 1234 aGV5dGhlcmU=", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			reqId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			sig, err := base64.StdEncoding.DecodeString(args[1])
			if err != nil {
				return err
			}

			msg := &v1beta3.MsgFulfilSignRequest{
				Creator:   clientCtx.GetFromAddress().String(),
				RequestId: reqId,
				Status:    v1beta3.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				Result: &v1beta3.MsgFulfilSignRequest_Payload{
					Payload: &v1beta3.MsgSignedData{
						SignedData: sig,
					},
				},
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func RejectSignRequestTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reject-sign-request [request-id] [reject-reason]",
		Short: "Reject a signature request providing a reason.",
		Long: `Reject a signature request providing a reason.
The sender of this transaction must be a writer of the Keychain for the request.`,
		Example: fmt.Sprintf("%s tx warden reject-sign-request 1234 oops", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			reqId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := &v1beta3.MsgFulfilSignRequest{
				Creator:   clientCtx.GetFromAddress().String(),
				RequestId: reqId,
				Status:    v1beta3.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				Result: &v1beta3.MsgFulfilSignRequest_RejectReason{
					RejectReason: args[1],
				},
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewInferenceRequestTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-inference-request --input [input]",
		Short: "Create a new inference request",
		Long: `Create a new inference request by providing an input.

The input is a list of numbers.
The first number is used to indicate how many vectors there will be.
The following numbers are the vectors.

E.g. the list:
	10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20
will be interpreted as 10 vectors of size 2:
	(1, 2), (3, 4), (5, 6), (7, 8), (9, 10), (11, 12), (13, 14), (15, 16), (17, 18), (19, 20)
`,
		Example: fmt.Sprintf("%s tx warden new-inference-request --input 10,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20", version.AppName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			input, err := cmd.Flags().GetFloat64Slice("input")
			if err != nil {
				return err
			}

			if len(input) < 2 {
				return fmt.Errorf("input must contain at least two numbers")
			}

			vectorsCount := input[0]
			if math.Trunc(vectorsCount) != vectorsCount {
				return fmt.Errorf("size of vectors must be an integer")
			}
			if vectorsCount <= 0 {
				return fmt.Errorf("size of vectors must be greater than 0")
			}

			if (len(input)-1)%int(vectorsCount) != 0 {
				return fmt.Errorf("expected %d vectors (got %d numbers)", int(vectorsCount), len(input)-1)
			}

			inputBz, err := inference.Input(input).Serialize()
			if err != nil {
				return err
			}

			callbackContractAddr, err := cmd.Flags().GetString("callback-contract-addr")
			if err != nil {
				return err
			}

			msg := &v1beta3.MsgNewInferenceRequest{
				Creator:          clientCtx.GetFromAddress().String(),
				Input:            inputBz,
				ContractCallback: callbackContractAddr,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Float64Slice("input", nil, "Input vectors")
	cmd.Flags().String("callback-contract-addr", "", "CosmWasm address of the contract to call after inference is done")

	return cmd
}
