package pubkey

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32/legacybech32" //nolint
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	"github.com/joltify-finance/joltify_lending/x/third_party/evmutil/types"
)

// PoolPubKeyToJoltAddress return the jolt encoded pubkey
func PoolPubKeyToJoltAddress(pk string) (sdk.AccAddress, error) {
	pubkey, err := legacybech32.UnmarshalPubKey(legacybech32.AccPK, pk) //nolint
	if err != nil {
		return sdk.AccAddress{}, err
	}
	addr, err := sdk.AccAddressFromHexUnsafe(pubkey.Address().String())
	return addr, err
}

func keyConvert() { //nolint
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("jolt", "joltpub")

	privKey := "460232de02330e492ae9afb7a16c3de4d3268a8310451d9ce842953ca9b448a0"
	data, err := hex.DecodeString(privKey)
	if err != nil {
		panic("fail to decode the sk")
	}
	sk := secp256k1.PrivKey{Key: data}

	sss := sk.PubKey().Address()
	fmt.Printf(">>>>ssss %v\n", sss.Bytes())

	accAddr, err := sdk.AccAddressFromHexUnsafe(sk.PubKey().Address().String())
	if err != nil {
		panic(err)
	}

	fmt.Printf(">>>Jolt address %v\n", accAddr.String())

	// eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee
	pkEth := ethsecp256k1.PrivKey{
		Key: data,
	}

	pk1 := pkEth.PubKey()

	pk2 := ethsecp256k1.PubKey{
		Key: pk1.Bytes(),
	}

	fmt.Printf(">>>jolt %v\n", sk.PubKey().Bytes())
	fmt.Printf(">>eth>>%v\n", pk2.Bytes())

	fmt.Printf("ttttttt>>>>%v\n", pk2.Address().Bytes())

	addrEth := common.BytesToAddress(pk2.Address())
	fmt.Printf(">>ethAddr>>>%v\n", addrEth.String())

	cosaddr := sdk.AccAddress(addrEth.Bytes())
	fmt.Printf(">>>convert from eth %v\n", cosaddr.String())

	frompub := sk.PubKey().Bytes()
	target := pkEth.PubKey().Bytes()

	fmt.Printf("%v\n%v\n", frompub, target)

	ok := bytes.Equal(frompub, target)
	fmt.Printf(">>>equal %v\n", ok)

	in := common.Bytes2Hex(sk.PubKey().Bytes())
	bb := common.Hex2Bytes(in)
	fmt.Printf(">>%v\n", bb)
	fmt.Printf(">>equal??>>%v\n", bytes.Equal(bb, sk.PubKey().Bytes()))
	fmt.Printf(">IIIIIIII>%v\n", in)
	eAddr, _ := types.PubKeyToEthAddr(in)

	coAddr, _ := types.PubKeyToJoltAddr(in)
	fmt.Printf(">>>%v\n", eAddr.String())
	fmt.Printf(">>>%v\n", coAddr.String())
}
