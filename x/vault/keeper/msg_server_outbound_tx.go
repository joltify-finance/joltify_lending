package keeper

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

func (k msgServer) sanitize(msg *types.MsgCreateOutboundTx) bool {
	msg.OutboundTx = strings.ToLower(msg.OutboundTx)
	msg.InTxHash = strings.ToLower(msg.InTxHash)
	data, err := hex.DecodeString(msg.InTxHash)
	if err != nil {
		return false
	}
	var needMintStr string
	if msg.NeedMint {
		needMintStr = "true"
	} else {
		needMintStr = "false"
	}
	// we need to ensure all the input are lower case
	target := crypto.Keccak256Hash(msg.ReceiverAddress.Bytes(), []byte(msg.ChainType), []byte(needMintStr), data)
	return strings.EqualFold(target.Hex(), msg.RequestID)
}

func (k msgServer) CreateOutboundTx(goCtx context.Context, msg *types.MsgCreateOutboundTx) (*types.MsgCreateOutboundTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sanitize(msg) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the request ID do not match the given inbound tx and receiver")
	}

	height, err := strconv.ParseInt(msg.BlockHeight, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid block height %v", msg.BlockHeight))
	}

	history, get := k.vaultStaking.GetHistoricalInfo(ctx, height)
	if !get {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("too early, we cannot find the block %v", msg.BlockHeight))
	}

	// now we check whether the msg is sent from the validator
	validators := history.GetValset()
	isValidator := false
	for _, el := range validators {
		if el.GetOperator().Equals(msg.Creator) {
			isValidator = true
			break
		}
	}
	if !isValidator {
		ctx.Logger().Info("not a validator update tss message", "result", "false")
		return &types.MsgCreateOutboundTxResponse{Successful: false}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintln("not a validator"))
	}
	info, isFound := k.GetOutboundTx(ctx, msg.RequestID)
	if isFound {
		proposal, ok := info.Items[msg.OutboundTx]
		if ok {
			for _, el := range proposal.Entry {
				if el.Address.Equals(msg.Creator) {
					ctx.Logger().Info("the creator has already submitted the outbound tx")
					return &types.MsgCreateOutboundTxResponse{Successful: false}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintln("already submitted"))
				}
			}
			thisProposal := types.Entity{Address: msg.Creator, Feecoin: msg.Feecoin}
			proposal.Entry = append(proposal.Entry, &thisProposal)
			info.Items[msg.OutboundTx] = proposal

			// here we check whether we need to send the fee to the validator module
			k.sendFeeToStakes(ctx, len(validators), msg.OutboundTx, &info)

			k.SetOutboundTx(
				ctx,
				info,
			)
			return &types.MsgCreateOutboundTxResponse{Successful: true}, nil
		}

		thisProposal := types.Entity{Address: msg.Creator, Feecoin: msg.Feecoin}
		info.Items[msg.OutboundTx] = types.Proposals{Entry: []*types.Entity{&thisProposal}}
		k.SetOutboundTx(
			ctx,
			info,
		)
		return &types.MsgCreateOutboundTxResponse{Successful: true}, nil
	}

	items := make(map[string]types.Proposals)

	thisProposal := types.Entity{Address: msg.Creator, Feecoin: msg.Feecoin}
	items[msg.OutboundTx] = types.Proposals{Entry: []*types.Entity{&thisProposal}}
	newInfo := types.OutboundTx{
		Index:           msg.RequestID,
		Items:           items,
		Processed:       false,
		ReceiverAddress: msg.ReceiverAddress,
		ChainType:       msg.ChainType,
		InTxHash:        msg.InTxHash,
		NeedMint:        msg.NeedMint,
	}

	k.SetOutboundTx(
		ctx,
		newInfo,
	)
	return &types.MsgCreateOutboundTxResponse{Successful: true}, nil
}

func (k msgServer) sendFeeToStakes(ctx sdk.Context, totalValidatorNum int, outboundTx string, info *types.OutboundTx) {
	if info.Processed {
		return
	}

	req := types.QueryLatestPoolRequest{}

	lastPoolsInfoResp, err := k.GetLastPool(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		ctx.Logger().Error("vault", "error", "fail to get the last pool")
		return
	}
	if len(lastPoolsInfoResp.Pools) < 2 {
		ctx.Logger().Error("vault", "error", "less than two pool we skip fee distribution")
		return
	}

	candidateDec := sdk.NewDecWithPrec(int64(totalValidatorNum), 0)
	params := k.GetParams(ctx)
	candidateNumDec := candidateDec.Mul(params.CandidateRatio).MulTruncate(sdk.MustNewDecFromStr("0.6667"))
	candidateNum := int(candidateNumDec.TruncateInt64())

	proposal := info.Items[outboundTx]
	if len(proposal.Entry) < candidateNum {
		return
	}
	feeCoinMap := make(map[string]int)
	for _, el := range proposal.Entry {
		feeCoinMap[el.Feecoin.String()]++
		if feeCoinMap[el.Feecoin.String()] >= candidateNum {
			// transfer
			previousFee := k.GetAllFeeAmount(ctx)
			previousFee.Sort()
			el.Feecoin.Sort()
			newFee := previousFee.Add(el.Feecoin...)
			k.SetStoreFeeAmount(ctx, newFee)
			info.Feecoin = el.Feecoin
			if info.NeedMint {
				err := k.bankKeeper.MintCoins(ctx, types.ModuleName, newFee)
				if err != nil {
					ctx.Logger().Error("vault", "fail to mint fee", el.Feecoin.String())
					return
				}
				err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lastPoolsInfoResp.Pools[0].CreatePool.PoolAddr, newFee)
				if err != nil {
					ctx.Logger().Error("vault", "fail to send fee to latest pool", el.Feecoin.String())
					return
				}
			}
			info.Processed = true
			ctx.Logger().Info("vault", "fee to be distributed", el.Feecoin.String())
			break
		}
	}
}
