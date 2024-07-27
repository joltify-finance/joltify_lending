package keeper

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	coserrors "cosmossdk.io/errors"
	"cosmossdk.io/x/nft"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func parameterSanitize(payFreqStr string, apyStr []string) ([]sdkmath.LegacyDec, int32, error) {
	apyJunior, err := sdkmath.LegacyNewDecFromStr(apyStr[0])
	if err != nil {
		return nil, 0, err
	}

	apySenior, err := sdkmath.LegacyNewDecFromStr(apyStr[1])
	if err != nil {
		return nil, 0, err
	}

	payFreq, err := strconv.ParseInt(payFreqStr, 10, 64)
	if err != nil {
		panic("incorrect payfreq format")
	}
	if payFreq > types.Maxfreq || payFreq < types.Minfreq {
		return nil, 0, errors.New("pay frequency is invalid")
	}
	return []sdkmath.LegacyDec{apyJunior, apySenior}, int32(payFreq), nil
}

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.TargetTokenAmount.IsZero() {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidVersion, "the amount cannot be 0")
	}

	targetProject, ok := k.kycKeeper.GetProject(ctx, msg.ProjectIndex)
	if !ok {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "the given project %v cannot be found", msg.ProjectIndex)
	}

	_, err := k.priceFeedKeeper.GetCurrentPrice(ctx, targetProject.MarketId)
	if err != nil {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "the given marketID %v cannot be found", targetProject.MarketId)
	}

	spvAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	if !targetProject.ProjectOwner.Equals(spvAddress) {
		return nil, coserrors.Wrapf(errorsmod.ErrUnauthorized, "unauthorized address %v", msg.Creator)
	}

	apys, payfreq, err := parameterSanitize(targetProject.PayFreq, msg.Apy)
	if err != nil {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "invalid parameter: %v", err.Error())
	}

	poolTypes := []string{types.Junior, types.Senior}
	indexHashResp := make([]string, 0, 2)

	pa := k.GetParams(ctx)

	var conversion int32

	conversion = -1

	for _, market := range pa.Markets {
		if market.GetDenom() == msg.TargetTokenAmount[0].Denom {
			conversion = market.GetConversionFactor()
			break
		}
	}

	if conversion < 0 {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "invalid parameter from market: %v", "conversion factor")
	}

	gap := 18 - conversion
	if gap < 0 {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "invalid parameter: %v", "conversion factor")
	}

	adj := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(gap)), nil)
	minBorrowAmount := targetProject.MinBorrowAmount.Quo(sdkmath.NewIntFromBigInt(adj))
	minDeposit := targetProject.MinDepositAmount.Quo(sdkmath.NewIntFromBigInt(adj))

	supportedTokens := pa.BurnThreshold

	var typePrefix string
	// sort the pool and returned otherwise the test may fail as it assume the pool comes with senior first
	for index, targetAmount := range msg.TargetTokenAmount {
		tcs := sdk.NewCoins(targetAmount)
		if !tcs.DenomsSubsetOf(supportedTokens) {
			return nil, coserrors.Wrapf(types.ErrInvalidParameter, "invalid parameter: %v", "unsupported token")
		}

		typePrefix = poolTypes[index]
		poolApy := apys[index]
		poolType := poolTypes[index]
		ePoolType := types.PoolInfo_SENIOR
		if poolType == types.Junior {
			ePoolType = types.PoolInfo_JUNIOR
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
			ProjectName:                   targetProject.BasicInfo.ProjectName,
			LinkedProject:                 msg.ProjectIndex,
			OwnerAddress:                  spvAddress,
			Apy:                           poolApy,
			TargetAmount:                  targetAmount,
			MinDepositAmount:              minDeposit,
			PayFreq:                       payfreq,
			ReserveFactor:                 types.RESERVEFACTOR,
			PoolNFTIds:                    []string{},
			PoolStatus:                    types.PoolInfo_PREPARE,
			PoolType:                      ePoolType,
			ProjectLength:                 targetProject.ProjectLength,
			LastPaymentTime:               ctx.BlockTime(),
			BorrowedAmount:                sdk.NewCoin(denomPrefix+targetAmount.Denom, sdkmath.NewInt(0)),
			UsableAmount:                  sdk.NewCoin(targetAmount.Denom, sdkmath.NewInt(0)),
			EscrowInterestAmount:          sdkmath.NewInt(0),
			EscrowPrincipalAmount:         sdk.NewCoin(targetAmount.Denom, sdkmath.NewInt(0)),
			WithdrawProposalAmount:        sdk.NewCoin(denomPrefix+targetAmount.Denom, sdkmath.NewInt(0)),
			WithdrawAccounts:              make([]sdk.AccAddress, 0, 200),
			TransferAccounts:              make([]sdk.AccAddress, 0, 200),
			ProcessedTransferAccounts:     make([]sdk.AccAddress, 0, 200), // this is used to track transferred accounts when we close the pool
			ProcessedWithdrawAccounts:     make([]sdk.AccAddress, 0, 200), // this is used to track the withdrawal accounts when we close the pool
			TotalTransferOwnershipAmount:  sdk.NewCoin(denomPrefix+targetAmount.Denom, sdkmath.ZeroInt()),
			MinBorrowAmount:               sdk.NewCoin(targetAmount.Denom, minBorrowAmount),
			WithdrawRequestWindowSeconds:  targetProject.WithdrawRequestWindowSeconds,
			PoolLockedSeconds:             targetProject.PoolLockedSeconds,
			PoolTotalBorrowLimit:          targetProject.PoolTotalBorrowLimit,
			CurrentPoolTotalBorrowCounter: 0,
			PoolCreatedTime:               ctx.BlockTime(),
			GraceTime:                     targetProject.GraceTime,
			PoolDenomPrefix:               denomPrefix,
			SeparatePool:                  targetProject.SeparatePool,
		}

		k.SetPool(ctx, poolInfo)
		err = k.NftKeeper.SaveClass(ctx, poolNFTClass)
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
