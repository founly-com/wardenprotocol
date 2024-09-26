package warden

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cmn "github.com/evmos/evmos/v18/precompiles/common"
	wardencommon "github.com/warden-protocol/wardenprotocol/precompiles/common"
	types "github.com/warden-protocol/wardenprotocol/warden/x/warden/types/v1beta3"
)

func newMsgAddKeychainAdmin(args []interface{}, origin common.Address) (*types.MsgAddKeychainAdminRequest, *common.Address, error) {
	if len(args) != 2 {
		return nil, nil, fmt.Errorf(cmn.ErrInvalidNumberOfArgs, 2, len(args))
	}

	keychainId := args[0].(uint64)
	newAdminAddress := args[1].(common.Address)
	authority := wardencommon.Bech32StrFromAddress(origin)
	newAdmin := wardencommon.Bech32StrFromAddress(newAdminAddress)

	return &types.MsgAddKeychainAdminRequest{
		Authority:  authority,
		KeychainId: keychainId,
		NewAdmin:   newAdmin,
	}, &newAdminAddress, nil
}

func newMsgAddKeychainWriter(args []interface{}, origin common.Address) (*types.MsgAddKeychainWriter, *common.Address, error) {
	if len(args) != 2 {
		return nil, nil, fmt.Errorf(cmn.ErrInvalidNumberOfArgs, 2, len(args))
	}

	keychainId := args[0].(uint64)
	newAdminAddress := args[1].(common.Address)
	creator := wardencommon.Bech32StrFromAddress(origin)
	newWriter := wardencommon.Bech32StrFromAddress(newAdminAddress)

	return &types.MsgAddKeychainWriter{
		Creator:    creator,
		KeychainId: keychainId,
		Writer:     newWriter,
	}, &newAdminAddress, nil
}

func newMsgFulfilKeyRequest(args []interface{}, keyRequestStatus types.KeyRequestStatus, origin common.Address) (*types.MsgFulfilKeyRequest, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf(cmn.ErrInvalidNumberOfArgs, 2, len(args))
	}

	creator := wardencommon.Bech32StrFromAddress(origin)
	requestId := args[0].(uint64)
	if keyRequestStatus == types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED {
		key := args[1].([]byte)
		result := &types.MsgFulfilKeyRequest_Key{
			Key: &types.MsgNewKey{
				PublicKey: key,
			},
		}

		return &types.MsgFulfilKeyRequest{
			Creator:   creator,
			RequestId: requestId,
			Status:    keyRequestStatus,
			Result:    result,
		}, nil
	} else {
		rejectReason := args[1].(string)
		result := &types.MsgFulfilKeyRequest_RejectReason{
			RejectReason: rejectReason,
		}

		return &types.MsgFulfilKeyRequest{
			Creator:   creator,
			RequestId: requestId,
			Status:    keyRequestStatus,
			Result:    result,
		}, nil
	}
}

func newMsgFulfilSignRequest(args []interface{}, signRequestStatus types.SignRequestStatus, origin common.Address) (*types.MsgFulfilSignRequest, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf(cmn.ErrInvalidNumberOfArgs, 2, len(args))
	}

	creator := wardencommon.Bech32StrFromAddress(origin)
	requestId := args[0].(uint64)
	if signRequestStatus == types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED {
		signedData := args[1].([]byte)
		result := &types.MsgFulfilSignRequest_Payload{
			Payload: &types.MsgSignedData{
				SignedData: signedData,
			},
		}

		return &types.MsgFulfilSignRequest{
			Creator:   creator,
			RequestId: requestId,
			Status:    types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
			Result:    result,
		}, nil
	} else {
		rejectReason := args[1].(string)
		result := &types.MsgFulfilSignRequest_RejectReason{
			RejectReason: rejectReason,
		}

		return &types.MsgFulfilSignRequest{
			Creator:   creator,
			RequestId: requestId,
			Status:    signRequestStatus,
			Result:    result,
		}, nil
	}
}