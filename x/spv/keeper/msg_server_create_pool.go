package keeper

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

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

func calculateApys(targetAmount, pool1Amount sdk.Coin, baseApy, pool1Apy sdk.Dec, isJunior bool) (map[string]sdk.Dec, map[string]sdk.Coin, error) {
	poolsInfoAPY := make(map[string]sdk.Dec)
	poolsInfoAmount := make(map[string]sdk.Coin)

	if targetAmount.IsLTE(pool1Amount) {
		return nil, nil, errors.New("amount incorrect")
	}

	pool2Amount := targetAmount.Sub(pool1Amount)
	if isJunior {
		poolsInfoAmount["junior"] = pool1Amount
		poolsInfoAmount["senior"] = pool2Amount
	} else {
		poolsInfoAmount["junior"] = pool2Amount
		poolsInfoAmount["senior"] = pool1Amount
	}

	if pool1Amount.Amount.LT(sdk.ZeroInt()) || pool2Amount.Amount.LT(sdk.ZeroInt()) {
		return nil, nil, errors.New("one pool has less than 0 amount")
	}

	//ij := sdk.NewDecFromInt(pool1Amount.Amount).Mul(pool1Apy)
	ij := pool1Apy.MulInt(pool1Amount.Amount)
	it := baseApy.MulInt(targetAmount.Amount)
	pool2Apy := it.Sub(ij).QuoTruncate(sdk.NewDecFromInt(pool2Amount.Amount))

	if isJunior {
		poolsInfoAPY["junior"] = pool1Apy
		poolsInfoAPY["senior"] = pool2Apy
	} else {
		poolsInfoAPY["junior"] = pool2Apy
		poolsInfoAPY["senior"] = pool1Apy
	}

	if pool1Apy.LT(sdk.ZeroDec()) || pool2Apy.LT(sdk.ZeroDec()) {
		return nil, nil, errors.New("one apy has less than 0 ")
	}

	return poolsInfoAPY, poolsInfoAmount, nil

}

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	allProjects := k.kycKeeper.GetProjects(ctx)

	if allProjects == nil || int32(len(allProjects)) < msg.ProjectIndex {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "the given project %v cannot be found", msg.ProjectIndex)
	}

	if msg.TargetTokenAmount.IsZero() {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidVersion, "the amount cannot be 0")
	}

	targetProject := allProjects[msg.ProjectIndex-1]
	_, err := k.priceFeedKeeper.GetCurrentPrice(ctx, targetProject.MarketId)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "the given marketID %v cannot be found", targetProject.MarketId)
	}

	if targetProject.ProjectTargetAmount.IsLT(msg.TargetTokenAmount) {
		return nil, coserrors.Wrap(sdkerrors.ErrInvalidRequest, "the junior amout is larger than the project target amount")
	}

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

	poolsInfoAPY, poolsInfoAmount, err := calculateApys(targetProject.ProjectTargetAmount, msg.TargetTokenAmount, targetProject.BaseApy, apy, true)

	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "junior pool amount larger than target")
	}

	var indexHashResp []string
	var typePrefix string
	for poolType, amount := range poolsInfoAmount {

		poolApy := poolsInfoAPY[poolType]

		enuPoolType := types.PoolInfo_JUNIOR
		if poolType == "senior" {
			enuPoolType = types.PoolInfo_SENIOR
			typePrefix = "senior"
		} else {
			typePrefix = "junior"
		}

		indexHash := crypto.Keccak256Hash([]byte(targetProject.BasicInfo.ProjectName), spvAddress.Bytes(), []byte(poolType))
		urlHash := crypto.Keccak256Hash([]byte(targetProject.BasicInfo.ProjectsUrl))

		indexHashResp = append(indexHashResp, indexHash.Hex())
		_, found := k.GetPools(ctx, indexHash.Hex())
		if found {
			return nil, coserrors.Wrapf(types.ErrPoolExisted, "pool existed")
		}

		nftClassID := fmt.Sprintf("class-%v", indexHash.String()[2:])
		poolNFTClass := nft.Class{
			Id:          nftClassID,
			Name:        msg.PoolName + "-" + typePrefix,
			Symbol:      "asset-" + indexHash.Hex(),
			Description: targetProject.BasicInfo.Description,
			Uri:         targetProject.BasicInfo.ProjectsUrl,
			UriHash:     urlHash.Hex(),
		}

		denomPrefix := strings.Split(targetProject.MarketId, ":")[0] + "-"
		poolInfo := types.PoolInfo{
			Index:                         indexHash.Hex(),
			PoolName:                      msg.PoolName + "-" + typePrefix,
			LinkedProject:                 msg.ProjectIndex,
			OwnerAddress:                  spvAddress,
			Apy:                           poolApy,
			TargetAmount:                  amount,
			PayFreq:                       payfreq,
			ReserveFactor:                 types.RESERVEFACTOR,
			PoolNFTIds:                    []string{},
			PoolStatus:                    types.PoolInfo_PREPARE,
			PoolType:                      enuPoolType,
			ProjectLength:                 targetProject.ProjectLength,
			LastPaymentTime:               ctx.BlockTime(),
			PrincipalPaid:                 false,
			BorrowedAmount:                sdk.NewCoin(denomPrefix+msg.TargetTokenAmount.Denom, sdk.NewInt(0)),
			UsableAmount:                  sdk.NewCoin(msg.TargetTokenAmount.Denom, sdk.NewInt(0)),
			EscrowInterestAmount:          sdk.NewInt(0),
			EscrowPrincipalAmount:         sdk.NewCoin(msg.TargetTokenAmount.Denom, sdk.NewInt(0)),
			WithdrawProposalAmount:        sdk.NewCoin(denomPrefix+msg.TargetTokenAmount.Denom, sdk.NewInt(0)),
			WithdrawAccounts:              make([]sdk.AccAddress, 0, 200),
			TransferAccounts:              make([]sdk.AccAddress, 0, 200),
			WithdrawRequestWindowSeconds:  targetProject.WithdrawRequestWindowSeconds,
			PoolLockedSeconds:             targetProject.PoolLockedSeconds,
			PoolTotalBorrowLimit:          targetProject.PoolTotalBorrowLimit,
			CurrentPoolTotalBorrowCounter: 0,
			PoolCreatedTime:               ctx.BlockTime(),
			GraceTime:                     targetProject.GraceTime,
			PoolDenomPrefix:               denomPrefix,
		}

		k.SetPool(ctx, poolInfo)
		err = k.nftKeeper.SaveClass(ctx, poolNFTClass)
		if err != nil {
			return nil, err
		}
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
		),
	)

	return &types.MsgCreatePoolResponse{PoolIndex: indexHashResp}, nil
}
