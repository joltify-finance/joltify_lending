package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/stretchr/testify/assert"
)

func TestProcessHistory(t *testing.T) {
	q := types.CoinsQuota{
		History:  []*types.HistoricalAmount{},
		CoinsSum: sdk.NewCoins(),
	}

	bnb := sdk.NewCoin("bnb", sdk.NewInt(10))
	busd := sdk.NewCoin("busd", sdk.NewInt(2))
	eth := sdk.NewCoin("eth", sdk.NewInt(4))

	bnb2 := sdk.NewCoin("bnb", sdk.NewInt(1))
	busd2 := sdk.NewCoin("busd", sdk.NewInt(2))
	eth2 := sdk.NewCoin("eth", sdk.NewInt(3))

	f1 := newHistory(1, sdk.Coins{})

	f2 := newHistory(2, sdk.NewCoins(bnb, busd))
	f3 := newHistory(3, sdk.NewCoins(bnb, busd))
	f4 := newHistory(4, sdk.NewCoins(bnb, eth))

	f5 := newHistory(1, sdk.Coins{})
	f6 := newHistory(6, sdk.NewCoins(bnb2, busd2))
	f7 := newHistory(7, sdk.NewCoins(bnb2))
	f8 := newHistory(8, sdk.NewCoins(bnb2, eth2))
	f9 := newHistory(9, sdk.NewCoins(eth2))
	f10 := newHistory(10, sdk.Coins{})

	processHistory(5, f1, &q)
	processHistory(5, f2, &q)
	processHistory(5, f3, &q)
	processHistory(5, f4, &q)
	processHistory(5, f5, &q)
	expected := sdk.NewCoins(sdk.NewCoin("bnb", sdk.NewInt(30)), sdk.NewCoin("busd", sdk.NewInt(4)), sdk.NewCoin("eth", sdk.NewInt(4)))
	assert.True(t, q.CoinsSum.IsEqual(expected))
	assert.True(t, q.History[0].Amount.Empty())
	assert.True(t, q.History[1].Amount.IsEqual(f2.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[4].Amount.Empty())

	processHistory(5, f6, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f2.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f6.Amount))

	expected = expected.Add(f6.Amount...)
	expected = expected.Sub(f1.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	processHistory(5, f7, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f7.Amount))

	expected = expected.Add(f7.Amount...)
	expected = expected.Sub(f2.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	processHistory(5, f8, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f8.Amount))

	expected = expected.Add(f8.Amount...)
	expected = expected.Sub(f3.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	processHistory(5, f9, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f8.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f9.Amount))

	expected = expected.Add(f9.Amount...)
	expected = expected.Sub(f4.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	processHistory(5, f10, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f8.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f9.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f10.Amount))
	expected = expected.Add(f10.Amount...)
	expected = expected.Sub(f5.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))
}

func TestProcessHistoryWithManyEmpty(t *testing.T) {
	q := types.CoinsQuota{
		History:  []*types.HistoricalAmount{},
		CoinsSum: sdk.NewCoins(),
	}

	bnb := sdk.NewCoin("bnb", sdk.NewInt(10))
	busd := sdk.NewCoin("busd", sdk.NewInt(2))
	eth := sdk.NewCoin("eth", sdk.NewInt(4))

	bnb2 := sdk.NewCoin("bnb", sdk.NewInt(1))
	busd2 := sdk.NewCoin("busd", sdk.NewInt(2))
	eth2 := sdk.NewCoin("eth", sdk.NewInt(3))

	f1 := newHistory(1, sdk.Coins{})

	f2 := newHistory(2, sdk.NewCoins(bnb, busd))
	f3 := newHistory(3, sdk.Coins{})
	f4 := newHistory(4, sdk.NewCoins(bnb, eth))

	f5 := newHistory(1, sdk.Coins{})
	f6 := newHistory(6, sdk.NewCoins(bnb2, busd2))
	f7 := newHistory(7, sdk.Coins{})
	f8 := newHistory(8, sdk.Coins{})
	f9 := newHistory(9, sdk.NewCoins(eth2))
	f10 := newHistory(10, sdk.Coins{})

	processHistory(5, f1, &q)
	processHistory(5, f2, &q)
	processHistory(5, f3, &q)
	processHistory(5, f4, &q)
	processHistory(5, f5, &q)
	expected := sdk.NewCoins(sdk.NewCoin("bnb", sdk.NewInt(20)), sdk.NewCoin("busd", sdk.NewInt(2)), sdk.NewCoin("eth", sdk.NewInt(4)))
	assert.True(t, q.CoinsSum.IsEqual(expected))
	assert.True(t, q.History[0].Amount.Empty())
	assert.True(t, q.History[1].Amount.IsEqual(f2.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[4].Amount.Empty())

	processHistory(5, f6, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f2.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f6.Amount))

	expected = expected.Add(f6.Amount...)
	expected = expected.Sub(f1.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	processHistory(5, f7, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f7.Amount))

	expected = expected.Add(f7.Amount...)
	expected = expected.Sub(f2.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	processHistory(5, f8, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f8.Amount))

	expected = expected.Add(f8.Amount...)
	expected = expected.Sub(f3.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	processHistory(5, f9, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f8.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f9.Amount))

	expected = expected.Add(f9.Amount...)
	expected = expected.Sub(f4.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	processHistory(5, f10, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f8.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f9.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f10.Amount))

	expected = expected.Add(f10.Amount...)
	expected = expected.Sub(f5.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	processHistory(5, f10, &q)
	processHistory(5, f10, &q)
	processHistory(5, f10, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f9.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f10.Amount))

	expected = expected.Add(f10.Amount...)
	expected = expected.Sub(f6.Amount)
	expected = expected.Sub(f7.Amount)
	expected = expected.Sub(f8.Amount)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	processHistory(5, f10, &q)
	processHistory(5, f10, &q)
	processHistory(5, f10, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f10.Amount))
	assert.True(t, q.CoinsSum.Empty())
}
