package keeper

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func parameterSanitize(msg *types.MsgCreatePool) (sdk.Dec, int32, error) {
	apy, err := sdk.NewDecFromStr(msg.Apy)
	if err != nil {
		return sdk.Dec{}, 0, err
	}
	payFreq, err := strconv.ParseInt(msg.PayFreq, 10, 64)
	if payFreq > types.Maxfreq || payFreq < types.Minfreq {
		return sdk.Dec{}, 0, errors.New("pay frequency is invalid")
	}
	return apy, int32(payFreq), nil
}

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	allProjects := k.kycKeeper.GetProjects(ctx)

	if allProjects == nil || int32(len(allProjects)) < msg.ProjectIndex {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "the given project %v cannot be found", msg.ProjectIndex)
	}

	targetProject := allProjects[msg.ProjectIndex-1]

	spvAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	if !targetProject.ProjectOwner.Equals(spvAddress) {
		return nil, coserrors.Wrapf(sdkerrors.ErrUnauthorized, "unauthorized address %v", msg.Creator)
	}

	apy, payfreq, err := parameterSanitize(msg)
	if err != nil {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "invalid parameter: %v", err.Error())
	}

	indexHash := crypto.Keccak256Hash([]byte(targetProject.BasicInfo.Description), spvAddress.Bytes(), apy.BigInt().Bytes())
	urlHash := crypto.Keccak256Hash([]byte(targetProject.BasicInfo.ProjectsUrl))

	_, found := k.GetPools(ctx, indexHash.Hex())
	if found {
		return nil, coserrors.Wrapf(types.ErrPoolExisted, "pool existed")
	}

	nftClassID := fmt.Sprintf("nft-%v", indexHash.String()[2:])
	poolNFTClass := nft.Class{
		Id:          nftClassID,
		Name:        msg.PoolName,
		Symbol:      "asset-" + indexHash.Hex(),
		Description: targetProject.BasicInfo.Description,
		Uri:         targetProject.BasicInfo.ProjectsUrl,
		UriHash:     urlHash.Hex(),
	}

	poolInfo := types.PoolInfo{
		Index:            indexHash.Hex(),
		PoolName:         msg.PoolName,
		LinkedProject:    msg.ProjectIndex,
		OwnerAddress:     spvAddress,
		Apy:              apy,
		TotalAmount:      msg.TargetTokenAmount,
		PayFreq:          payfreq,
		ReserveFactor:    types.RESERVEFACTOR,
		PoolNFTIds:       []string{},
		PoolStatus:       types.PoolInfo_INACTIVE,
		ProjectLength:    targetProject.ProjectLength,
		BorrowedAmount:   sdk.NewCoin(msg.TargetTokenAmount.Denom, sdk.NewInt(0)),
		BorrowableAmount: sdk.NewCoin(msg.TargetTokenAmount.Denom, sdk.NewInt(0)),
	}

	k.SetPool(ctx, poolInfo)
	k.nftKeeper.SaveClass(ctx, poolNFTClass)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
		),
	)

	return &types.MsgCreatePoolResponse{}, nil
}
