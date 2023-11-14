package keeper

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	abci "github.com/cometbft/cometbft/abci/types"
)

func doQueryAllLiquidate(ctx sdk.Context, k Keeper, reqBorrowers sdk.AccAddress) ([]types2.LiquidateItem, error) {
	var liquidateUsers []types2.LiquidateItem
	var borrows types2.Borrows

	if !reqBorrowers.Empty() {
		b := types2.Borrow{
			Borrower: reqBorrowers,
		}
		borrows = append(borrows, b)
	} else {
		k.IterateBorrows(ctx, func(borrow types2.Borrow) (stop bool) {
			borrows = append(borrows, borrow)
			return false
		})
	}

	var syncedBorrows types2.Borrows
	for _, borrow := range borrows {
		syncedBorrow, _ := k.GetSyncedBorrow(ctx, borrow.Borrower)
		syncedBorrows = append(syncedBorrows, syncedBorrow)
	}

	var syncedDeposit types2.Deposits
	for _, el := range syncedBorrows {
		deposit, found := k.GetSyncedDeposit(ctx, el.Borrower)
		if !found {
			return nil, types2.ErrDepositNotFound
		}
		syncedDeposit = append(syncedDeposit, deposit)
	}
	for i := range syncedBorrows {
		eachBorrow := syncedBorrows[i]
		eachDeposit := syncedDeposit[i]
		// ratio= borrow / deposit
		_, ratio, err := k.IsWithinValidLtvRange(ctx, eachDeposit, eachBorrow)
		if err != nil {
			return nil, err
		}
		if ratio.Equal(sdk.MustNewDecFromStr("0")) || ratio.GTE(sdk.MustNewDecFromStr("0.95")) {
			users := types2.LiquidateItem{
				Owner: eachBorrow.Borrower.String(),
				Ltv:   ratio.String(),
			}
			liquidateUsers = append(liquidateUsers, users)
		}
	}
	return liquidateUsers, nil
}

func queryGetLiquidate(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var liquidateReq types2.QueryLiquidate
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &liquidateReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// we query all the borrows
	var retLiquidateResp []types2.LiquidateItem
	var bz []byte
	if liquidateReq.Borrow == "" {
		ret, err := doQueryAllLiquidate(ctx, k, sdk.AccAddress{})
		if err != nil {
			return nil, sdkerrors.Wrap(errors.New("err in query the liquidate users"), err.Error())
		}

		start, end := client.Paginate(len(ret), liquidateReq.Page, liquidateReq.Limit, 100)
		if start < 0 || end < 0 {
			bz, err = codec.MarshalJSONIndent(legacyQuerierCdc, retLiquidateResp)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}
			return bz, nil
		}
		retLiquidateResp = ret[start:end]
		bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, retLiquidateResp)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return bz, nil
	}

	v, err := sdk.AccAddressFromBech32(liquidateReq.Borrow)
	if err != nil {
		return nil, err
	}

	ret, err := doQueryAllLiquidate(ctx, k, v)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.New("err in query the liquidate users"), err.Error())
	}

	start, end := client.Paginate(len(ret), liquidateReq.Page, liquidateReq.Limit, 100)
	if start < 0 || end < 0 {
		bz, err = codec.MarshalJSONIndent(legacyQuerierCdc, retLiquidateResp)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return bz, nil
	}
	retLiquidateResp = ret[start:end]
	bz, err = codec.MarshalJSONIndent(legacyQuerierCdc, retLiquidateResp)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
