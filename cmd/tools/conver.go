package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
)

func main() {
	inAddr := os.Args[1]
	AddrBytes, err := sdk.GetFromBech32(inAddr, "jolt")
	if err != nil {
		fmt.Printf("invalid address format")
		return
	}

	fmt.Printf(">>>>%v\n", AddrBytes)
	ethAddr := common.BytesToAddress(AddrBytes)
	fmt.Printf("%v\n", ethAddr.String())

	privKey := "64285613d3569bcaa7a24c9d74d4cec5c18dcf6a08e4c0f66596078f3a4a35b5"

	data, err := hex.DecodeString(privKey)
	if err != nil {
		panic("fail to decode the sk")
	}
	sk := secp256k1.PrivKey{Key: data}
	//pk, err := legacybech32.MarshalPubKey(legacybech32.AccPK, sk.PubKey())
	//if err != nil {
	//	panic(err)
	//}

	ethPrivkey := ethsecp256k1.PrivKey{
		Key: data,
	}

	ethAddr = common.BytesToAddress(ethPrivkey.PubKey().Address().Bytes())
	fmt.Printf(">>>>eth addr%v\n", ethAddr.String())

	config := sdk.GetConfig()
	config.SetCoinType(60)
	config.SetBech32PrefixForAccount("jolt", "joltpub")
	config.SetBech32PrefixForValidator("joltval", "joltvpub")
	config.SetBech32PrefixForConsensusNode("joltvalcons", "joltcpub")

	addr, err := sdk.AccAddressFromHexUnsafe(sk.PubKey().Address().String())
	if err != nil {
		panic("err")
	}

	fmt.Printf(">>>eth pubket %v\n", ethPrivkey.PubKey().String())
	fmt.Printf(">>>cos pubket %v\n", sk.PubKey().String())

	//addr2, err := sdk.AccAddressFromHexUnsafe(pubkey.Address().String())
	//if err != nil {
	//	panic("err")
	//}

	fmt.Printf(">>>>recover %v\n", addr.String())

}
