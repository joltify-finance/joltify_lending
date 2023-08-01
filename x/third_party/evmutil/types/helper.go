package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
)

func PubKeyToEthAddr(pubKey string) (common.Address, error) {
	pkBytes := common.Hex2Bytes(pubKey)
	pk := ethsecp256k1.PubKey{
		Key: pkBytes,
	}

	if len(pk.Key) != secp256k1.PubKeySize {
		return common.Address{}, fmt.Errorf("invalid pubkey(%v) length", pubKey)
	}
	addrEth := common.BytesToAddress(pk.Address())
	return addrEth, nil
}

func PubKeyToJoltAddr(pubKey string) (sdk.AccAddress, error) {
	pkBytes := common.Hex2Bytes(pubKey)
	pk := secp256k1.PubKey{
		Key: pkBytes,
	}

	if len(pk.Key) != secp256k1.PubKeySize {
		return nil, fmt.Errorf("invalid pubkey(%v) length", pubKey)
	}

	accAddr, err := sdk.AccAddressFromHexUnsafe(pk.Address().String())
	if err != nil {
		return nil, err
	}
	return accAddr, nil
}

func EVMAddressToJoltAddress(addr common.Address) sdk.AccAddress {
	accAddr := sdk.AccAddress(addr.Bytes())
	return accAddr
}
