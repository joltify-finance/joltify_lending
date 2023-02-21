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

func parameterSanitize(payFreqStr, apyStr string) (sdk.Dec, int32, error) {
	apy, err := sdk.NewDecFromStr(apyStr)
	if err != nil {
		return sdk.Dec{}, 0, err
	}
	payFreq, err := strconv.ParseInt(payFreqStr, 10, 64)
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

	apy, payfreq, err := parameterSanitize(targetProject.PayFreq, msg.Apy)
	if err != nil {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "invalid parameter: %v", err.Error())
	}

	poolsInfoAPY := make(map[string]sdk.Dec)
	poolsInfoAmount := make(map[string]sdk.Coin)

	if targetProject.ProjectTargetAmount.IsLTE(msg.TargetTokenAmount) {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "junior pool amount larger thatn target")
	}

	seniorAmount := targetProject.ProjectTargetAmount.Sub(msg.TargetTokenAmount)
	poolsInfoAmount["junior"] = msg.TargetTokenAmount
	poolsInfoAmount["senior"] = seniorAmount

	poolsInfoAPY["junior"] = apy

	ij := sdk.NewDecFromInt(msg.TargetTokenAmount.Amount).Mul(apy)
	it := sdk.NewDecFromInt(targetProject.ProjectTargetAmount.Amount).Mul(targetProject.BaseApy)
	apySenior := it.Sub(ij).Quo(sdk.NewDecFromInt(seniorAmount.Amount))

	poolsInfoAPY["junior"] = apy
	poolsInfoAPY["senior"] = apySenior

	for poolType, amount := range poolsInfoAmount {

		poolApy := poolsInfoAPY[poolType]

		indexHash := crypto.Keccak256Hash([]byte(targetProject.BasicInfo.ProjectName), spvAddress.Bytes(), []byte(poolType))
		urlHash := crypto.Keccak256Hash([]byte(targetProject.BasicInfo.ProjectsUrl))

		fmt.Printf(">>>%v\n", indexHash)

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
			Apy:              poolApy,
			TargetAmount:     amount,
			PayFreq:          payfreq,
			ReserveFactor:    types.RESERVEFACTOR,
			PoolNFTIds:       []string{},
			PoolStatus:       types.PoolInfo_PREPARE,
			ProjectLength:    targetProject.ProjectLength,
			BorrowedAmount:   sdk.NewCoin(msg.TargetTokenAmount.Denom, sdk.NewInt(0)),
			BorrowableAmount: sdk.NewCoin(msg.TargetTokenAmount.Denom, sdk.NewInt(0)),
			TotalAmount:      sdk.NewCoin(msg.TargetTokenAmount.Denom, sdk.NewInt(0)),
		}

		k.SetPool(ctx, poolInfo)
		k.nftKeeper.SaveClass(ctx, poolNFTClass)
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
		),
	)

	return &types.MsgCreatePoolResponse{}, nil
}
