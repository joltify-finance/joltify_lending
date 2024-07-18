package ocr

import (
	"os"
	"runtime/debug"

	"github.com/rs/zerolog"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/joltify-finance/joltify_lending/x/third_party/ocr/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/ocr/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (res *sdk.Result, err error) {
		defer Recover(&err)

		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *types.MsgCreateFeed:
			res, err := msgServer.CreateFeed(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateFeed:
			res, err := msgServer.UpdateFeed(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgTransmit:
			res, err := msgServer.Transmit(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgFundFeedRewardPool:
			res, err := msgServer.FundFeedRewardPool(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgWithdrawFeedRewardPool:
			res, err := msgServer.WithdrawFeedRewardPool(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetPayees:
			res, err := msgServer.SetPayees(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgTransferPayeeship:
			res, err := msgServer.TransferPayeeship(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAcceptPayeeship:
			res, err := msgServer.AcceptPayeeship(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateParams:
			res, err := msgServer.UpdateParams(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, errors.Wrapf(sdkerrors.ErrUnknownRequest, "Unrecognized OCR msg type: %T", msg)
		}
	}
}

func Recover(err *error) { // nolint:all
	// fixme we do not need to have the recover!!
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			logger.Error().Err(e).Msg("ocr msg handler panicked with an error")
			logger.Info().Msg(string(debug.Stack()))
		} else {
			logger.Err(r.(error))
		}
	}
}
