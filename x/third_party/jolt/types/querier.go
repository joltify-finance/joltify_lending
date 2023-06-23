package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// Querier routes for the jolt module
const (
	QueryGetParams           = "params"
	QueryGetModuleAccounts   = "accounts"
	QueryGetDeposits         = "deposits"
	QueryGetUnsyncedDeposits = "unsynced-deposits"
	QueryGetTotalDeposited   = "total-deposited"
	QueryGetBorrows          = "borrows"
	QueryGetUnsyncedBorrows  = "unsynced-borrows"
	QueryGetTotalBorrowed    = "total-borrowed"
	QueryGetInterestRate     = "interest-rate"
	QueryGetReserves         = "reserves"
	QueryGetInterestFactors  = "interest-factors"
	QueryGetLiquidate        = "liquidate"
)

// QueryDepositsParams is the params for a filtered deposit query
type QueryDepositsParams struct {
	Page  int            `json:"page" yaml:"page"`
	Limit int            `json:"limit" yaml:"limit"`
	Denom string         `json:"denom" yaml:"denom"`
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
}

// QueryLiquidate is the params for a filtered deposit query
type QueryLiquidate struct {
	Page   int    `json:"page" yaml:"page"`
	Limit  int    `json:"limit" yaml:"limit"`
	Borrow string `json:"borrow" yaml:"borrow"`
}

// QueryUnsyncedDepositsParams is the params for a filtered unsynced deposit query.
type QueryUnsyncedDepositsParams struct {
	Page  int            `json:"page" yaml:"page"`
	Limit int            `json:"limit" yaml:"limit"`
	Denom string         `json:"denom" yaml:"denom"`
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
}

// QueryAccountParams is the params for a filtered module account query
type QueryAccountParams struct {
	Page  int    `json:"page" yaml:"page"`
	Limit int    `json:"limit" yaml:"limit"`
	Name  string `json:"name" yaml:"name"`
}

// ModAccountWithCoins includes the module account with its coins
type ModAccountWithCoins struct {
	Account authtypes.ModuleAccountI `json:"account" yaml:"account"`
	Coins   sdk.Coins                `json:"coins" yaml:"coins"`
}

// QueryBorrowsParams is the params for a filtered borrows query
type QueryBorrowsParams struct {
	Page  int            `json:"page" yaml:"page"`
	Limit int            `json:"limit" yaml:"limit"`
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Denom string         `json:"denom" yaml:"denom"`
}

// QueryUnsyncedBorrowsParams is the params for a filtered unsynced borrows query
type QueryUnsyncedBorrowsParams struct {
	Page  int            `json:"page" yaml:"page"`
	Limit int            `json:"limit" yaml:"limit"`
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Denom string         `json:"denom" yaml:"denom"`
}

// QueryTotalBorrowedParams is the params for a filtered total borrowed coins query
type QueryTotalBorrowedParams struct {
	Denom string `json:"denom" yaml:"denom"`
}

// QueryTotalDepositedParams is the params for a filtered total deposited coins query
type QueryTotalDepositedParams struct {
	Denom string `json:"denom" yaml:"denom"`
}

// QueryInterestRateParams is the params for a filtered interest rate query
type QueryInterestRateParams struct {
	Denom string `json:"denom" yaml:"denom"`
}

// MoneyMarketInterestRates is a slice of MoneyMarketInterestRate
type MoneyMarketInterestRates []MoneyMarketInterestRate

// QueryReservesParams is the params for a filtered reserves query
type QueryReservesParams struct {
	Denom string `json:"denom" yaml:"denom"`
}

// QueryInterestFactorsParams is the params for a filtered interest factors query
type QueryInterestFactorsParams struct {
	Denom string `json:"denom" yaml:"denom"`
}

// InterestFactors is a slice of InterestFactor
type InterestFactors []InterestFactor
