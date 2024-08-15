package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/app/config"
)

//func main() {
//	encoded := os.Args[1]
//	decoded, err := base64.StdEncoding.DecodeString(encoded)
//	if err != nil {
//		panic(err)
//	}
//
//	pk := secp256k1.PubKey{
//		Key: decoded,
//	}
//
//	accAddr, err := sdk.AccAddressFromHexUnsafe(pk.Address().String())
//	if err != nil {
//		return
//	}
//
//	accAddraa, err := bech32.ConvertAndEncode("jolt", accAddr.Bytes())
//	if err != nil {
//		return
//	}
//
//	pkAddr := common.Bytes2Hex(pk.Bytes())
//
//	fmt.Printf(">>>%v\n", accAddraa)
//	fmt.Printf(">>>%v\n", pkAddr)
//}

func main() {
	// get address
	config.SetupConfig()
	_, addr := app.GeneratePrivKeyAddressPairs(5)

	for _, el := range addr {
		valaddr := types.ValAddress(el.Bytes())
		fmt.Printf(">>>>%v\n", el.String(), valaddr.String())

	}
}
