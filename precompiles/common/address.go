package common

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// Bech32StrFromAddress Creates bech32 address string from eth address
func Bech32StrFromAddress(address common.Address) string {
	return sdk.AccAddress(address.Bytes()).String()
}

// AddressFromBech32Str Creates eth address from bech32 address string
func AddressFromBech32Str(address string) (common.Address, error) {
	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return common.Address{}, err
	}

	return common.BytesToAddress(accAddress.Bytes()), nil
}

// AddressesFromBech32StrItemArray tries to create a slice of common.Address values from arbitrary slice
func AddressesFromBech32StrItemArray[T any](items []T, addressFunc func(T) string) ([]common.Address, error) {
	ethAddresses := make([]common.Address, 0)

	for _, item := range items {
		ethAddress, err := AddressFromBech32Str(addressFunc(item))
		if err != nil {
			return nil, err
		}

		ethAddresses = append(ethAddresses, ethAddress)
	}

	return ethAddresses, nil
}

// AddressesFromBech32StrArray tries to create a slice of common.Address values from a string slice
func AddressesFromBech32StrArray(items []string) ([]common.Address, error) {
	id := func(i string) string {
		return i
	}

	return AddressesFromBech32StrItemArray(items, id)
}

// MustAddressFromBech32Str creates common.Address from a string, panics if it's unsuccessful
func MustAddressFromBech32Str(address string) common.Address {
	accAddress, err := AddressFromBech32Str(address)
	if err != nil {
		panic(err)
	}

	return accAddress
}