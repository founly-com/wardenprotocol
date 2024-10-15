package act

import (
	"bytes"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	cmn "github.com/evmos/evmos/v20/precompiles/common"
)

const (
	// EventCreateTemplate defines the event type for the x/act CreateTemplate transaction.
	EventCreateTemplate = "CreateTemplate"
	// EventUpdateTemplate defines the event type for the x/act UpdateTemplate transaction.
	EventUpdateTemplate = "UpdateTemplate"
	// EventActionVoted defines the event type for the x/act VoteForAction transaction.
	EventActionVoted = "ActionVoted"
	// EventCreateAction defines the event type for the x/act NewAction transaction.
	EventCreateAction      = "CreateAction"
	EventActionStateChange = "ActionStateChange"
)

func (p *Precompile) GetCreateTemplateEvent(ctx sdk.Context, writerAddress *ethcmn.Address, _ sdk.Msg) (*ethtypes.Log, error) {
	return p.getEthEvent(
		ctx,
		*writerAddress,
		EventCreateTemplate,
		parseNewTemplateEvent)
}

func parseNewTemplateEvent(sdkEvent sdk.Event) (*bytes.Buffer, error) {
	var b bytes.Buffer

	for _, attr := range sdkEvent.Attributes {
		key := attr.GetKey()
		val := strings.Trim(attr.GetValue(), "\"")
		switch key {
		case "id":
			keychainId, success := new(big.Int).SetString(val, 10)
			if !success {
				return nil, fmt.Errorf("NewTemplateEvent: invalid keychain id type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(keychainId)))

		case "creator":
			address, err := sdk.AccAddressFromBech32(val)
			if err != nil {
				return nil, fmt.Errorf("NewTemplateEvent: invalid writers count type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(ethcmn.Address(address.Bytes()))))
		}
	}

	return &b, nil
}

func (p *Precompile) GetUpdateTemplateEvent(ctx sdk.Context, writerAddress ethcmn.Address) (*ethtypes.Log, error) {
	return p.getEthEvent(
		ctx,
		writerAddress,
		EventUpdateTemplate,
		parseUpdateTemplateEvent)
}

func parseUpdateTemplateEvent(sdkEvent sdk.Event) (*bytes.Buffer, error) {
	var b bytes.Buffer

	for _, attr := range sdkEvent.Attributes {
		key := attr.GetKey()
		val := strings.Trim(attr.GetValue(), "\"")
		switch key {
		case "id":
			keychainId, success := new(big.Int).SetString(val, 10)
			if !success {
				return nil, fmt.Errorf("UpdateTemplateEvent: invalid id type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(keychainId)))
		}

	}

	return &b, nil
}

func (p *Precompile) GetCreateActionEvent(ctx sdk.Context, writerAddress ethcmn.Address) (*ethtypes.Log, error) {
	return p.getEthEvent(
		ctx,
		writerAddress,
		EventCreateAction,
		parseCreateActionEvent)
}

func parseCreateActionEvent(sdkEvent sdk.Event) (*bytes.Buffer, error) {
	var b bytes.Buffer

	for _, attr := range sdkEvent.Attributes {
		key := attr.GetKey()
		val := strings.Trim(attr.GetValue(), "\"")
		switch key {
		case "id":
			keychainId, success := new(big.Int).SetString(val, 10)
			if !success {
				return nil, fmt.Errorf("CreateActionEvent: invalid id type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(keychainId)))
		case "creator":
			address, err := sdk.AccAddressFromBech32(val)
			if err != nil {
				return nil, fmt.Errorf("CreateActionEvent: invalid creator type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(ethcmn.Address(address.Bytes()))))
		}

	}
	return &b, nil
}

func (p *Precompile) GetActionVotedEvent(ctx sdk.Context, writerAddress *ethcmn.Address, _ sdk.Msg) (*ethtypes.Log, error) {
	return p.getEthEvent(
		ctx,
		*writerAddress,
		EventActionVoted,
		parseActionVotedEvent)
}

func parseActionVotedEvent(sdkEvent sdk.Event) (*bytes.Buffer, error) {
	var b bytes.Buffer

	for _, attr := range sdkEvent.Attributes {
		key := attr.GetKey()
		val := strings.Trim(attr.GetValue(), "\"")
		switch key {
		case "id":
			keychainId, success := new(big.Int).SetString(val, 10)
			if !success {
				return nil, fmt.Errorf("ActionVotedEvent: invalid keychain id type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(keychainId)))
		case "participant":
			address, err := sdk.AccAddressFromBech32(val)
			if err != nil {
				return nil, fmt.Errorf("ActionVotedEvent: invalid participant type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(ethcmn.Address(address.Bytes()))))

		case "vote_type":
			value, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("ActionVotedEvent: invalid vote_type type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(value)))
		}

	}
	return &b, nil
}

func (p *Precompile) GetActionStateChangeEvent(ctx sdk.Context, writerAddress *ethcmn.Address, _ sdk.Msg) (*ethtypes.Log, error) {
	return p.getEthEvent(
		ctx,
		*writerAddress,
		EventActionStateChange,
		parseActionStateChangeEvent)
}

func parseActionStateChangeEvent(sdkEvent sdk.Event) (*bytes.Buffer, error) {
	var b bytes.Buffer

	for _, attr := range sdkEvent.Attributes {
		key := attr.GetKey()
		val := strings.Trim(attr.GetValue(), "\"")
		switch key {
		case "id":
			value, success := new(big.Int).SetString(val, 10)
			if !success {
				return nil, fmt.Errorf("ActionStateChangeEvent: invalid id type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(value)))
		case "previous_status":
			value, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("ActionStateChangeEvent: invalid previous_status type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(value)))
		case "new_status":
			value, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("ActionStateChangeEvent: invalid new_status type")
			}

			b.Write(cmn.PackNum(reflect.ValueOf(value)))
		}

	}

	return &b, nil
}

func (p *Precompile) getEthEvent(
	ctx sdk.Context,
	writerAddress ethcmn.Address,
	eventName string,
	eventParser func(sdk.Event) (*bytes.Buffer, error)) (*ethtypes.Log, error) {

	// Prepare the event topics
	event := p.ABI.Events[eventName]
	sdkEvents := ctx.EventManager().Events()
	sdkEventLen := len(sdkEvents)
	for i := range sdkEvents {
		// iterage in reverse order
		x := sdkEvents[sdkEventLen-1-i]
		// TODO: check type, contract that .sol Event name should match with sdk Event name
		fmt.Printf("\nx.Type %v, eventName %v\n", x.Type, eventName)
		if x.Type == eventName {
			b, err := eventParser(x)
			if err != nil {
				return nil, err
			}

			topics := make([]ethcmn.Hash, 2)
			// The first topic is always the signature of the event.
			topics[0] = event.ID

			topics[1], err = cmn.MakeTopic(writerAddress)
			if err != nil {
				return nil, err
			}

			ethLog := ethtypes.Log{
				Address:     p.Address(),
				Topics:      topics,
				Data:        b.Bytes(),
				BlockNumber: uint64(ctx.BlockHeight()),
			}

			return &ethLog, nil
		}
	}

	return nil, nil
}