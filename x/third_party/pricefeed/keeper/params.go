package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"
)

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx context.Context) types.Params {
	var p types.Params
	fmt.Printf(">>>>>>>>>>>>>%v\n", p.ParamSetPairs())
	k.paramSubspace.GetParamSet(sdk.UnwrapSDKContext(ctx), &p)
	return p
}

// SetParams sets params on the store
func (k Keeper) SetParams(ctx context.Context, params types.Params) {
	k.paramSubspace.SetParamSet(sdk.UnwrapSDKContext(ctx), &params)
}

// GetMarkets returns the markets from params
func (k Keeper) GetMarkets(ctx context.Context) types.Markets {
	return k.GetParams(ctx).Markets
}

// GetOracles returns the oracles in the pricefeed store
func (k Keeper) GetOracles(ctx context.Context, marketID string) ([]sdk.AccAddress, error) {
	for _, m := range k.GetMarkets(ctx) {
		if marketID == m.MarketID {
			return m.Oracles, nil
		}
	}
	return nil, errorsmod.Wrap(types.ErrInvalidMarket, marketID)
}

// GetOracle returns the oracle from the store or an error if not found
func (k Keeper) GetOracle(ctx context.Context, marketID string, address sdk.AccAddress) (sdk.AccAddress, error) {
	oracles, err := k.GetOracles(ctx, marketID)
	if err != nil {
		// Error already wrapped
		return nil, err
	}
	for _, addr := range oracles {
		if addr.Equals(address) {
			return addr, nil
		}
	}
	return nil, errorsmod.Wrap(types.ErrInvalidOracle, address.String())
}

// GetMarket returns the market if it is in the pricefeed system
func (k Keeper) GetMarket(ctx context.Context, marketID string) (types.Market, bool) {
	markets := k.GetMarkets(ctx)

	for i := range markets {
		if markets[i].MarketID == marketID {
			return markets[i], true
		}
	}
	return types.Market{}, false
}

// GetAuthorizedAddresses returns a list of addresses that have special authorization within this module, eg the oracles of all markets.
func (k Keeper) GetAuthorizedAddresses(ctx context.Context) []sdk.AccAddress {
	var oracles []sdk.AccAddress
	uniqueOracles := map[string]bool{}

	for _, m := range k.GetMarkets(ctx) {
		for _, o := range m.Oracles {
			// de-dup list of oracles
			if _, found := uniqueOracles[o.String()]; !found {
				oracles = append(oracles, o)
			}
			uniqueOracles[o.String()] = true
		}
	}
	return oracles
}
