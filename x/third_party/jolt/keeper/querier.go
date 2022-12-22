package keeper

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types2.QueryGetParams:
			return queryGetParams(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetModuleAccounts:
			return queryGetModAccounts(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetDeposits:
			return queryGetDeposits(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetUnsyncedDeposits:
			return queryGetUnsyncedDeposits(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetTotalDeposited:
			return queryGetTotalDeposited(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetBorrows:
			return queryGetBorrows(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetUnsyncedBorrows:
			return queryGetUnsyncedBorrows(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetTotalBorrowed:
			return queryGetTotalBorrowed(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetInterestRate:
			return queryGetInterestRate(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetReserves:
			return queryGetReserves(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetInterestFactors:
			return queryGetInterestFactors(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetLiquidate:
			return queryGetLiquidate(ctx, req, k, legacyQuerierCdc)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint", types2.ModuleName)
		}
	}
}

func queryGetParams(ctx sdk.Context, _ abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	// Get params
	params := k.GetParams(ctx)

	// Encode results
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryGetModAccounts(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryAccountParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	var accs []authtypes.ModuleAccountI
	if len(params.Name) > 0 {
		acc := k.accountKeeper.GetModuleAccount(ctx, params.Name)
		accs = append(accs, acc)
	} else {
		acc := k.accountKeeper.GetModuleAccount(ctx, types2.ModuleAccountName)
		accs = append(accs, acc)
	}

	// Include module account coins with its account to keep backwards compatibility with v39 account behavior
	response := make([]types2.ModAccountWithCoins, len(accs))
	for i, acc := range accs {
		coins := k.bankKeeper.GetAllBalances(ctx, acc.GetAddress())
		response[i] = types2.ModAccountWithCoins{
			Account: acc,
			Coins:   coins,
		}
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, response)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryGetDeposits(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryDepositsParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	denom := len(params.Denom) > 0
	owner := len(params.Owner) > 0

	var deposits types2.Deposits
	switch {
	case owner && denom:
		deposit, found := k.GetSyncedDeposit(ctx, params.Owner)
		if found {
			for _, coin := range deposit.Amount {
				if coin.Denom == params.Denom {
					deposits = append(deposits, deposit)
				}
			}
		}
	case owner:
		deposit, found := k.GetSyncedDeposit(ctx, params.Owner)
		if found {
			deposits = append(deposits, deposit)
		}
	case denom:
		k.IterateDeposits(ctx, func(deposit types2.Deposit) (stop bool) {
			if deposit.Amount.AmountOf(params.Denom).IsPositive() {
				deposits = append(deposits, deposit)
			}
			return false
		})
	default:
		k.IterateDeposits(ctx, func(deposit types2.Deposit) (stop bool) {
			deposits = append(deposits, deposit)
			return false
		})
	}

	var bz []byte

	// If owner param was specified then deposits array already contains the user's synced deposit
	if owner {
		bz, err = codec.MarshalJSONIndent(legacyQuerierCdc, deposits)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return bz, nil
	}

	// Otherwise we need to simulate syncing of each deposit
	var syncedDeposits types2.Deposits
	for _, deposit := range deposits {
		syncedDeposit, _ := k.GetSyncedDeposit(ctx, deposit.Depositor)
		syncedDeposits = append(syncedDeposits, syncedDeposit)
	}

	start, end := client.Paginate(len(syncedDeposits), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		syncedDeposits = types2.Deposits{}
	} else {
		syncedDeposits = syncedDeposits[start:end]
	}

	bz, err = codec.MarshalJSONIndent(legacyQuerierCdc, syncedDeposits)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryGetUnsyncedDeposits(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryUnsyncedDepositsParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	denom := len(params.Denom) > 0
	owner := len(params.Owner) > 0

	var deposits types2.Deposits
	switch {
	case owner && denom:
		deposit, found := k.GetDeposit(ctx, params.Owner)
		if found {
			for _, coin := range deposit.Amount {
				if coin.Denom == params.Denom {
					deposits = append(deposits, deposit)
				}
			}
		}
	case owner:
		deposit, found := k.GetDeposit(ctx, params.Owner)
		if found {
			deposits = append(deposits, deposit)
		}
	case denom:
		k.IterateDeposits(ctx, func(deposit types2.Deposit) (stop bool) {
			if deposit.Amount.AmountOf(params.Denom).IsPositive() {
				deposits = append(deposits, deposit)
			}
			return false
		})
	default:
		k.IterateDeposits(ctx, func(deposit types2.Deposit) (stop bool) {
			deposits = append(deposits, deposit)
			return false
		})
	}

	var bz []byte

	start, end := client.Paginate(len(deposits), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		deposits = types2.Deposits{}
	} else {
		deposits = deposits[start:end]
	}

	bz, err = codec.MarshalJSONIndent(legacyQuerierCdc, deposits)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryGetBorrows(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryBorrowsParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	denom := len(params.Denom) > 0
	owner := len(params.Owner) > 0

	var borrows types2.Borrows
	switch {
	case owner && denom:
		borrow, found := k.GetSyncedBorrow(ctx, params.Owner)
		if found {
			for _, coin := range borrow.Amount {
				if coin.Denom == params.Denom {
					borrows = append(borrows, borrow)
				}
			}
		}
	case owner:
		borrow, found := k.GetSyncedBorrow(ctx, params.Owner)
		if found {
			borrows = append(borrows, borrow)
		}
	case denom:
		k.IterateBorrows(ctx, func(borrow types2.Borrow) (stop bool) {
			if borrow.Amount.AmountOf(params.Denom).IsPositive() {
				borrows = append(borrows, borrow)
			}
			return false
		})
	default:
		k.IterateBorrows(ctx, func(borrow types2.Borrow) (stop bool) {
			borrows = append(borrows, borrow)
			return false
		})
	}

	var bz []byte

	// If owner param was specified then borrows array already contains the user's synced borrow
	if owner {
		bz, err = codec.MarshalJSONIndent(legacyQuerierCdc, borrows)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return bz, nil
	}

	// Otherwise we need to simulate syncing of each borrow
	var syncedBorrows types2.Borrows
	for _, borrow := range borrows {
		syncedBorrow, _ := k.GetSyncedBorrow(ctx, borrow.Borrower)
		syncedBorrows = append(syncedBorrows, syncedBorrow)
	}

	start, end := client.Paginate(len(syncedBorrows), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		syncedBorrows = types2.Borrows{}
	} else {
		syncedBorrows = syncedBorrows[start:end]
	}

	bz, err = codec.MarshalJSONIndent(legacyQuerierCdc, syncedBorrows)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryGetUnsyncedBorrows(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryUnsyncedBorrowsParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	denom := len(params.Denom) > 0
	owner := len(params.Owner) > 0

	var borrows types2.Borrows
	switch {
	case owner && denom:
		borrow, found := k.GetBorrow(ctx, params.Owner)
		if found {
			for _, coin := range borrow.Amount {
				if coin.Denom == params.Denom {
					borrows = append(borrows, borrow)
				}
			}
		}
	case owner:
		borrow, found := k.GetBorrow(ctx, params.Owner)
		if found {
			borrows = append(borrows, borrow)
		}
	case denom:
		k.IterateBorrows(ctx, func(borrow types2.Borrow) (stop bool) {
			if borrow.Amount.AmountOf(params.Denom).IsPositive() {
				borrows = append(borrows, borrow)
			}
			return false
		})
	default:
		k.IterateBorrows(ctx, func(borrow types2.Borrow) (stop bool) {
			borrows = append(borrows, borrow)
			return false
		})
	}

	var bz []byte

	start, end := client.Paginate(len(borrows), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		borrows = types2.Borrows{}
	} else {
		borrows = borrows[start:end]
	}

	bz, err = codec.MarshalJSONIndent(legacyQuerierCdc, borrows)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryGetTotalBorrowed(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryTotalBorrowedParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	borrowedCoins, found := k.GetBorrowedCoins(ctx)
	if !found {
		return nil, types2.ErrBorrowedCoinsNotFound
	}

	// If user specified a denom only return coins of that denom type
	if len(params.Denom) > 0 {
		borrowedCoins = sdk.NewCoins(sdk.NewCoin(params.Denom, borrowedCoins.AmountOf(params.Denom)))
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, borrowedCoins)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryGetTotalDeposited(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryTotalDepositedParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	suppliedCoins, found := k.GetSuppliedCoins(ctx)
	if !found {
		return nil, types2.ErrSuppliedCoinsNotFound
	}

	// If user specified a denom only return coins of that denom type
	if len(params.Denom) > 0 {
		suppliedCoins = sdk.NewCoins(sdk.NewCoin(params.Denom, suppliedCoins.AmountOf(params.Denom)))
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, suppliedCoins)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryGetInterestRate(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryInterestRateParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var moneyMarketInterestRates types2.MoneyMarketInterestRates
	var moneyMarkets types2.MoneyMarkets
	if len(params.Denom) > 0 {
		moneyMarket, found := k.GetMoneyMarket(ctx, params.Denom)
		if !found {
			return nil, types2.ErrMoneyMarketNotFound
		}
		moneyMarkets = append(moneyMarkets, moneyMarket)
	} else {
		moneyMarkets = k.GetAllMoneyMarkets(ctx)
	}

	// Calculate the borrow and supply APY interest rates for each money market
	for _, moneyMarket := range moneyMarkets {
		denom := moneyMarket.Denom
		macc := k.accountKeeper.GetModuleAccount(ctx, types2.ModuleName)
		cash := k.bankKeeper.GetBalance(ctx, macc.GetAddress(), denom).Amount

		borrowed := sdk.NewCoin(denom, sdk.ZeroInt())
		borrowedCoins, foundBorrowedCoins := k.GetBorrowedCoins(ctx)
		if foundBorrowedCoins {
			borrowed = sdk.NewCoin(denom, borrowedCoins.AmountOf(denom))
		}

		reserves, foundReserves := k.GetTotalReserves(ctx)
		if !foundReserves {
			reserves = sdk.NewCoins()
		}

		// CalculateBorrowRate calculates the current interest rate based on utilization (the fraction of supply that has been borrowed)
		borrowAPY, err := CalculateBorrowRate(moneyMarket.InterestRateModel, sdk.NewDecFromInt(cash), sdk.NewDecFromInt(borrowed.Amount), sdk.NewDecFromInt(reserves.AmountOf(denom)))
		if err != nil {
			return nil, err
		}

		utilRatio := CalculateUtilizationRatio(sdk.NewDecFromInt(cash), sdk.NewDecFromInt(borrowed.Amount), sdk.NewDecFromInt(reserves.AmountOf(denom)))
		fullSupplyAPY := borrowAPY.Mul(utilRatio)
		realSupplyAPY := fullSupplyAPY.Mul(sdk.OneDec().Sub(moneyMarket.ReserveFactor))

		moneyMarketInterestRate := types2.MoneyMarketInterestRate{
			Denom:              denom,
			SupplyInterestRate: realSupplyAPY.String(),
			BorrowInterestRate: borrowAPY.String(),
		}

		moneyMarketInterestRates = append(moneyMarketInterestRates, moneyMarketInterestRate)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, moneyMarketInterestRates)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryGetReserves(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryReservesParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	reserveCoins, found := k.GetTotalReserves(ctx)
	if !found {
		reserveCoins = sdk.Coins{}
	}

	// If user specified a denom only return coins of that denom type
	if len(params.Denom) > 0 {
		reserveCoins = sdk.NewCoins(sdk.NewCoin(params.Denom, reserveCoins.AmountOf(params.Denom)))
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, reserveCoins)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryGetInterestFactors(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryInterestFactorsParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var interestFactors types2.InterestFactors
	if len(params.Denom) > 0 {
		// Fetch supply/borrow interest factors for a single denom
		interestFactor := types2.InterestFactor{}
		interestFactor.Denom = params.Denom
		supplyInterestFactor, found := k.GetSupplyInterestFactor(ctx, params.Denom)
		if found {
			interestFactor.SupplyInterestFactor = supplyInterestFactor.String()
		}
		borrowInterestFactor, found := k.GetBorrowInterestFactor(ctx, params.Denom)
		if found {
			interestFactor.BorrowInterestFactor = borrowInterestFactor.String()
		}
		interestFactors = append(interestFactors, interestFactor)
	} else {
		interestFactorMap := make(map[string]types2.InterestFactor)
		// Populate mapping with supply interest factors
		k.IterateSupplyInterestFactors(ctx, func(denom string, factor sdk.Dec) (stop bool) {
			interestFactor := types2.InterestFactor{Denom: denom, SupplyInterestFactor: factor.String()}
			interestFactorMap[denom] = interestFactor
			return false
		})
		// Populate mapping with borrow interest factors
		k.IterateBorrowInterestFactors(ctx, func(denom string, factor sdk.Dec) (stop bool) {
			interestFactor, ok := interestFactorMap[denom]
			if !ok {
				newInterestFactor := types2.InterestFactor{Denom: denom, BorrowInterestFactor: factor.String()}
				interestFactorMap[denom] = newInterestFactor
			} else {
				interestFactor.BorrowInterestFactor = factor.String()
				interestFactorMap[denom] = interestFactor
			}
			return false
		})
		// Translate mapping to slice
		for _, val := range interestFactorMap {
			interestFactors = append(interestFactors, val)
		}
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, interestFactors)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

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
