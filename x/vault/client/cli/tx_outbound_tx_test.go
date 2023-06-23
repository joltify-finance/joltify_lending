package cli_test

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	app2 "github.com/joltify-finance/joltify_lending/app"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/joltify-finance/joltify_lending/x/vault/client/cli"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCreateOutboundTx(t *testing.T) {
	app2.SetSDKConfig()

	net, _ := preparePool(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	var fields []string
	for _, tc := range []struct {
		desc         string
		idRequestID  string
		blockHeight  string
		tx           string
		chainType    string
		inTxHash     string
		receiverAddr string
		args         []string
		err          error
		code         uint32
	}{
		{
			idRequestID:  strconv.Itoa(0),
			tx:           "testtoken",
			blockHeight:  "15",
			chainType:    "ETH",
			inTxHash:     "0a087f76c417d0284b54b49abc2952a0f13c9e9c1d6fffbdc6b6d89ed62f585f",
			receiverAddr: "jolt18mdnq8x9m07dryymlyf8jknagp87yga0hpe7n6",
			desc:         "valid",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			},
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			data, err := hex.DecodeString(tc.inTxHash)
			require.NoError(t, err)

			a, err := sdk.AccAddressFromBech32(tc.receiverAddr)
			require.NoError(t, err)

			target := crypto.Keccak256Hash(a.Bytes(), []byte(tc.chainType), []byte("false"), data)
			tc.idRequestID = target.Hex()

			args := []string{
				tc.idRequestID,
				tc.tx,
				tc.blockHeight,
				tc.chainType,
				tc.inTxHash,
				tc.receiverAddr,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			_, err = net.WaitForHeightWithTimeout(15, time.Minute)
			require.NoError(t, err)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateOutboundTx(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}
