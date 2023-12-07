package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	encoded := os.Args[1]
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}

	pk := secp256k1.PubKey{
		Key: decoded,
	}

	accAddr, err := sdk.AccAddressFromHexUnsafe(pk.Address().String())
	if err != nil {
		return
	}

	accAddraa, err := bech32.ConvertAndEncode("jolt", accAddr.Bytes())
	if err != nil {
		return
	}

	pkAddr := common.Bytes2Hex(pk.Bytes())

	fmt.Printf(">>>%v\n", accAddraa)
	fmt.Printf(">>>%v\n", pkAddr)
}
