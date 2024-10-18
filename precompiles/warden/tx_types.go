package warden

import (
	"fmt"

	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmosTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	wardencommon "github.com/warden-protocol/wardenprotocol/precompiles/common"
	actTypes "github.com/warden-protocol/wardenprotocol/warden/x/act/types/v1beta1"
	types "github.com/warden-protocol/wardenprotocol/warden/x/warden/types/v1beta3"
)

func newMsgAddKeychainAdmin(args []interface{}, origin common.Address) (*types.MsgAddKeychainAdminRequest, *common.Address, error) {
	if len(args) != 2 {
		return nil, nil, wardencommon.WrongArgsNumber{Expected: 2, Got: len(args)}
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
		return nil, nil, wardencommon.WrongArgsNumber{Expected: 2, Got: len(args)}
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
		return nil, wardencommon.WrongArgsNumber{Expected: 2, Got: len(args)}
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
		return nil, wardencommon.WrongArgsNumber{Expected: 2, Got: len(args)}
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

func newMsgNewKeychain(method *abi.Method, args []interface{}, origin common.Address) (*types.MsgNewKeychain, error) {
	if len(args) != 5 {
		return nil, wardencommon.WrongArgsNumber{Expected: 5, Got: len(args)}
	}

	creator := wardencommon.Bech32StrFromAddress(origin)
	var input newKeyChainInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to newKeyChainInput struct: %s", err)
	}

	return &types.MsgNewKeychain{
		Creator:      creator,
		Name:         input.Name,
		KeychainFees: input.KeychainFees,
		Description:  input.Description,
		Url:          input.Url,
		KeybaseId:    input.KeybaseId,
	}, nil
}

type newKeyChainInput struct {
	Name         string
	KeychainFees types.KeychainFees
	Description  string
	Url          string
	KeybaseId    string
}

func newMsgNewSpace(args []interface{}, origin common.Address) (*types.MsgNewSpace, error) {
	if len(args) != 5 {
		return nil, wardencommon.WrongArgsNumber{Expected: 5, Got: len(args)}
	}

	creator := wardencommon.Bech32StrFromAddress(origin)

	approveAdminTemplateId := args[0].(uint64)
	rejectAdminTemplateId := args[1].(uint64)
	approveSignTemplateId := args[2].(uint64)
	rejectSignTemplateId := args[3].(uint64)
	var additionalOwners []string
	for _, a := range args[4].([]common.Address) {
		additionalOwners = append(additionalOwners, wardencommon.Bech32StrFromAddress(a))
	}

	return &types.MsgNewSpace{
		Creator:                creator,
		ApproveAdminTemplateId: approveAdminTemplateId,
		RejectAdminTemplateId:  rejectAdminTemplateId,
		ApproveSignTemplateId:  approveSignTemplateId,
		RejectSignTemplateId:   rejectSignTemplateId,
		AdditionalOwners:       additionalOwners,
	}, nil
}

func newMsgRemoveKeychainAdmin(args []interface{}, origin common.Address) (*types.MsgRemoveKeychainAdminRequest, *common.Address, error) {
	if len(args) != 2 {
		return nil, nil, wardencommon.WrongArgsNumber{Expected: 2, Got: len(args)}
	}

	creator := wardencommon.Bech32StrFromAddress(origin)

	keychainId := args[0].(uint64)
	admin := args[1].(common.Address)

	return &types.MsgRemoveKeychainAdminRequest{
		Authority:  creator,
		KeychainId: keychainId,
		Admin:      wardencommon.Bech32StrFromAddress(admin),
	}, &admin, nil
}

func newMsgUpdateKeychain(method *abi.Method, args []interface{}, origin common.Address) (*types.MsgUpdateKeychain, error) {
	if len(args) != 6 {
		return nil, wardencommon.WrongArgsNumber{Expected: 6, Got: len(args)}
	}

	creator := wardencommon.Bech32StrFromAddress(origin)
	var input updateKeyChainInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to updateKeyChainInput struct: %s", err)
	}

	return &types.MsgUpdateKeychain{
		Creator:      creator,
		KeychainId:   input.KeychainId,
		Name:         input.Name,
		KeychainFees: mapEthKeychainFees(input.KeychainFees),
		Description:  input.Description,
		Url:          input.Url,
		KeybaseId:    input.KeybaseId,
	}, nil
}

type updateKeyChainInput struct {
	KeychainId   uint64
	Name         string
	KeychainFees KeychainFees
	Description  string
	Url          string
	KeybaseId    string
}

func newMsgAddSpaceOwner(args []interface{}, origin common.Address, act string) (*actTypes.MsgNewAction, error) {
	if len(args) != 6 {
		return nil, wardencommon.WrongArgsNumber{Expected: 6, Got: len(args)}
	}

	spaceId := args[0].(uint64)
	newOwnerAddress := args[1].(common.Address)
	nonce := args[2].(uint64)
	actionTimeoutHeight := args[3].(uint64)
	expectedApproveExpression := args[4].(string)
	expectedRejectExpression := args[5].(string)

	authority := wardencommon.Bech32StrFromAddress(origin)
	newOwner := wardencommon.Bech32StrFromAddress(newOwnerAddress)

	msgAddSpaceOwner := types.MsgAddSpaceOwner{
		Authority: act,
		SpaceId:   spaceId,
		NewOwner:  newOwner,
		Nonce:     nonce,
	}

	anyMsg, err := codecTypes.NewAnyWithValue(&msgAddSpaceOwner)
	if err != nil {
		return nil, err
	}

	return &actTypes.MsgNewAction{
		Creator:                   authority,
		Message:                   anyMsg,
		ActionTimeoutHeight:       actionTimeoutHeight,
		ExpectedApproveExpression: expectedApproveExpression,
		ExpectedRejectExpression:  expectedRejectExpression,
	}, nil
}

func newMsgRemoveSpaceOwner(args []interface{}, origin common.Address, act string) (*actTypes.MsgNewAction, error) {
	if len(args) != 6 {
		return nil, wardencommon.WrongArgsNumber{Expected: 6, Got: len(args)}
	}

	spaceId := args[0].(uint64)
	ownerAddress := args[1].(common.Address)
	nonce := args[2].(uint64)
	actionTimeoutHeight := args[3].(uint64)
	expectedApproveExpression := args[4].(string)
	expectedRejectExpression := args[5].(string)

	authority := wardencommon.Bech32StrFromAddress(origin)
	owner := wardencommon.Bech32StrFromAddress(ownerAddress)

	msgRemoveSpaceOwner := types.MsgRemoveSpaceOwner{
		Authority: act,
		SpaceId:   spaceId,
		Owner:     owner,
		Nonce:     nonce,
	}

	anyMsg, err := codecTypes.NewAnyWithValue(&msgRemoveSpaceOwner)
	if err != nil {
		return nil, err
	}

	return &actTypes.MsgNewAction{
		Creator:                   authority,
		Message:                   anyMsg,
		ActionTimeoutHeight:       actionTimeoutHeight,
		ExpectedApproveExpression: expectedApproveExpression,
		ExpectedRejectExpression:  expectedRejectExpression,
	}, nil
}

type newKeyRequestInput struct {
	SpaceId                   uint64
	KeychainId                uint64
	KeyType                   uint8
	ApproveTemplateId         uint64
	RejectTemplateId          uint64
	MaxKeychainFees           []cosmosTypes.Coin
	Nonce                     uint64
	ActionTimeoutHeight       uint64
	ExpectedApproveExpression string
	ExpectedRejectExpression  string
}

func newMsgNewKeyRequest(method *abi.Method, args []interface{}, origin common.Address, act string) (*actTypes.MsgNewAction, error) {
	if len(args) != 10 {
		return nil, wardencommon.WrongArgsNumber{Expected: 10, Got: len(args)}
	}

	var input newKeyRequestInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to newMsgNewKeyRequest struct: %s", err)
	}

	authority := wardencommon.Bech32StrFromAddress(origin)

	mapKeyType := func(keyType uint8) (types.KeyType, error) {
		switch keyType {
		case uint8(types.KeyType_KEY_TYPE_UNSPECIFIED):
			return types.KeyType_KEY_TYPE_UNSPECIFIED, nil
		case uint8(types.KeyType_KEY_TYPE_ECDSA_SECP256K1):
			return types.KeyType_KEY_TYPE_ECDSA_SECP256K1, nil
		case uint8(types.KeyType_KEY_TYPE_EDDSA_ED25519):
			return types.KeyType_KEY_TYPE_EDDSA_ED25519, nil
		default:
			return -1, fmt.Errorf("key type is not supported: %v", keyType)
		}
	}

	keyType, err := mapKeyType(input.KeyType)

	msgNewKeyRequest := types.MsgNewKeyRequest{
		Authority:         act,
		SpaceId:           input.SpaceId,
		KeychainId:        input.KeychainId,
		KeyType:           keyType,
		ApproveTemplateId: input.ApproveTemplateId,
		RejectTemplateId:  input.RejectTemplateId,
		MaxKeychainFees:   input.MaxKeychainFees,
		Nonce:             input.Nonce,
	}

	anyMsg, err := codecTypes.NewAnyWithValue(&msgNewKeyRequest)
	if err != nil {
		return nil, err
	}

	return &actTypes.MsgNewAction{
		Creator:                   authority,
		Message:                   anyMsg,
		ActionTimeoutHeight:       input.ActionTimeoutHeight,
		ExpectedApproveExpression: input.ExpectedApproveExpression,
		ExpectedRejectExpression:  input.ExpectedRejectExpression,
	}, nil
}

type newSignRequestInput struct {
	KeyId                     uint64
	Input                     []byte
	Analyzers                 []common.Address
	EncryptionKey             []byte
	MaxKeychainFees           []cosmosTypes.Coin
	Nonce                     uint64
	ActionTimeoutHeight       uint64
	ExpectedApproveExpression string
	ExpectedRejectExpression  string
}

func newMsgNewSignRequest(method *abi.Method, args []interface{}, origin common.Address, act string) (*actTypes.MsgNewAction, error) {
	if len(args) != 9 {
		return nil, wardencommon.WrongArgsNumber{Expected: 9, Got: len(args)}
	}

	var input newSignRequestInput
	if err := method.Inputs.Copy(&input, args); err != nil {
		return nil, fmt.Errorf("error while unpacking args to newSignRequestInput struct: %s", err)
	}

	var analyzers []string
	for _, a := range input.Analyzers {
		analyzers = append(analyzers, wardencommon.Bech32StrFromAddress(a))
	}

	authority := wardencommon.Bech32StrFromAddress(origin)

	msgNewSignRequest := types.MsgNewSignRequest{
		Authority:       act,
		KeyId:           input.KeyId,
		Input:           input.Input,
		Analyzers:       analyzers,
		EncryptionKey:   input.EncryptionKey,
		MaxKeychainFees: input.MaxKeychainFees,
		Nonce:           input.Nonce,
	}

	anyMsg, err := codecTypes.NewAnyWithValue(&msgNewSignRequest)
	if err != nil {
		return nil, err
	}

	return &actTypes.MsgNewAction{
		Creator:                   authority,
		Message:                   anyMsg,
		ActionTimeoutHeight:       input.ActionTimeoutHeight,
		ExpectedApproveExpression: input.ExpectedApproveExpression,
		ExpectedRejectExpression:  input.ExpectedRejectExpression,
	}, nil
}

func newMsgUpdateKey(args []interface{}, origin common.Address, act string) (*actTypes.MsgNewAction, error) {
	if len(args) != 6 {
		return nil, wardencommon.WrongArgsNumber{Expected: 6, Got: len(args)}
	}

	keyId := args[0].(uint64)
	approveTemplateId := args[1].(uint64)
	rejectTemplateId := args[2].(uint64)
	actionTimeoutHeight := args[3].(uint64)
	expectedApproveExpression := args[4].(string)
	expectedRejectExpression := args[5].(string)

	authority := wardencommon.Bech32StrFromAddress(origin)

	msgUpdateKey := types.MsgUpdateKey{
		Authority:         act,
		KeyId:             keyId,
		ApproveTemplateId: approveTemplateId,
		RejectTemplateId:  rejectTemplateId,
	}

	anyMsg, err := codecTypes.NewAnyWithValue(&msgUpdateKey)
	if err != nil {
		return nil, err
	}

	return &actTypes.MsgNewAction{
		Creator:                   authority,
		Message:                   anyMsg,
		ActionTimeoutHeight:       actionTimeoutHeight,
		ExpectedApproveExpression: expectedApproveExpression,
		ExpectedRejectExpression:  expectedRejectExpression,
	}, nil
}

func newMsgUpdateSpace(args []interface{}, origin common.Address, act string) (*actTypes.MsgNewAction, error) {
	if len(args) != 9 {
		return nil, wardencommon.WrongArgsNumber{Expected: 9, Got: len(args)}
	}

	spaceId := args[0].(uint64)
	nonce := args[1].(uint64)
	approveAdminTemplateId := args[2].(uint64)
	rejectAdminTemplateId := args[3].(uint64)
	approveSignTemplateId := args[4].(uint64)
	rejectSignTemplateId := args[5].(uint64)
	actionTimeoutHeight := args[6].(uint64)
	expectedApproveExpression := args[7].(string)
	expectedRejectExpression := args[8].(string)

	authority := wardencommon.Bech32StrFromAddress(origin)

	msgUpdateSpace := types.MsgUpdateSpace{
		Authority:              act,
		SpaceId:                spaceId,
		Nonce:                  nonce,
		ApproveAdminTemplateId: approveAdminTemplateId,
		RejectAdminTemplateId:  rejectAdminTemplateId,
		ApproveSignTemplateId:  approveSignTemplateId,
		RejectSignTemplateId:   rejectSignTemplateId,
	}

	anyMsg, err := codecTypes.NewAnyWithValue(&msgUpdateSpace)
	if err != nil {
		return nil, err
	}

	return &actTypes.MsgNewAction{
		Creator:                   authority,
		Message:                   anyMsg,
		ActionTimeoutHeight:       actionTimeoutHeight,
		ExpectedApproveExpression: expectedApproveExpression,
		ExpectedRejectExpression:  expectedRejectExpression,
	}, nil
}
