package tx

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/joltify-finance/joltify_lending/dydx_helper/module"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/constants"
)

// Returns the encoded msg as transaction. Will panic if encoding fails.
func MustGetTxBytes(msgs ...sdk.Msg) []byte {
	tx := constants.TestEncodingCfg.TxConfig.NewTxBuilder()
	err := tx.SetMsgs(msgs...)
	if err != nil {
		panic(err)
	}
	bz, err := constants.TestEncodingCfg.TxConfig.TxEncoder()(tx.GetTx())
	if err != nil {
		panic(err)
	}
	return bz
}

// Returns the account address that should sign the msg. Will panic if it is an unsupported message type.
func MustGetOnlySignerAddress(cdc codec.Codec, msg sdk.Msg) string {
	signers, _, err := cdc.GetMsgV1Signers(msg)
	if err != nil {
		panic(err)
	}
	if len(signers) == 0 {
		panic(fmt.Errorf("msg does not have designated signer: %T", msg))
	}
	if len(signers) > 1 {
		panic(fmt.Errorf("not supported - msg has multiple signers: %T", msg))
	}
	signer, err := module.InterfaceRegistry.SigningContext().AddressCodec().BytesToString(signers[0])
	if err != nil {
		panic(err)
	}

	v, err := module.InterfaceRegistry.SigningContext().AddressCodec().StringToBytes(signer)
	if err != nil {
		panic(err)
	}

	ad := sdk.AccAddress(v)
	return ad.String()
}
