package types

import (
	"errors"
	"fmt"
	"strings"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	JoltLiquidityProviderClaimType = "jolt_liquidity_provider"
	SwapClaimType                  = "swap"
	SPVClaimType                   = "spv"
)

// GetOwner is a getter for Claim Owner
func (c BaseClaim) GetOwner() sdk.AccAddress { return c.Owner }

// GetReward is a getter for Claim Reward
func (c BaseClaim) GetReward() sdk.Coin { return c.Reward }

// GetType returns the claim type, used to identify auctions in event attributes
func (c BaseClaim) GetType() string { return "base" }

// Validate performs a basic check of a BaseClaim fields
func (c BaseClaim) Validate() error {
	if c.Owner.Empty() {
		return errors.New("claim owner cannot be empty")
	}
	if !c.Reward.IsValid() {
		return fmt.Errorf("invalid reward amount: %s", c.Reward)
	}
	return nil
}

// GetOwner is a getter for Claim Owner
func (c BaseMultiClaim) GetOwner() sdk.AccAddress { return c.Owner }

// GetReward is a getter for Claim Reward
func (c BaseMultiClaim) GetReward() sdk.Coins { return c.Reward }

// GetType returns the claim type, used to identify auctions in event attributes
func (c BaseMultiClaim) GetType() string { return "base" }

// Validate performs a basic check of a BaseClaim fields
func (c BaseMultiClaim) Validate() error {
	if c.Owner.Empty() {
		return errors.New("claim owner cannot be empty")
	}
	if !c.Reward.IsValid() {
		return fmt.Errorf("invalid reward amount: %s", c.Reward)
	}
	return nil
}

// NewJoltLiquidityProviderClaim returns a new HardLiquidityProviderClaim
func NewJoltLiquidityProviderClaim(owner sdk.AccAddress, rewards sdk.Coins,
	supplyRewardIndexes, borrowRewardIndexes MultiRewardIndexes,
) JoltLiquidityProviderClaim {
	return JoltLiquidityProviderClaim{
		BaseMultiClaim: BaseMultiClaim{
			Owner:  owner,
			Reward: rewards,
		},
		SupplyRewardIndexes: supplyRewardIndexes,
		BorrowRewardIndexes: borrowRewardIndexes,
	}
}

// GetType returns the claim's type
func (j JoltLiquidityProviderClaim) GetType() string { return JoltLiquidityProviderClaimType }

// GetReward returns the claim's reward coin
func (j JoltLiquidityProviderClaim) GetReward() sdk.Coins { return j.Reward }

// GetOwner returns the claim's owner
func (j JoltLiquidityProviderClaim) GetOwner() sdk.AccAddress { return j.Owner }

// Validate performs a basic check of a HardLiquidityProviderClaim fields
func (j JoltLiquidityProviderClaim) Validate() error {
	if err := j.SupplyRewardIndexes.Validate(); err != nil {
		return err
	}

	if err := j.BorrowRewardIndexes.Validate(); err != nil {
		return err
	}

	return j.BaseMultiClaim.Validate()
}

// HasSupplyRewardIndex check if a claim has a supply reward index for the input collateral type
func (j JoltLiquidityProviderClaim) HasSupplyRewardIndex(denom string) (int64, bool) {
	for index, ri := range j.SupplyRewardIndexes {
		if ri.CollateralType == denom {
			return int64(index), true
		}
	}
	return 0, false
}

// HasBorrowRewardIndex check if a claim has a borrow reward index for the input collateral type
func (j JoltLiquidityProviderClaim) HasBorrowRewardIndex(denom string) (int64, bool) {
	for index, ri := range j.BorrowRewardIndexes {
		if ri.CollateralType == denom {
			return int64(index), true
		}
	}
	return 0, false
}

// JoltLiquidityProviderClaims slice of HardLiquidityProviderClaim
type JoltLiquidityProviderClaims []JoltLiquidityProviderClaim

// Validate checks if all the claims are valid and there are no duplicated
// entries.
func (cs JoltLiquidityProviderClaims) Validate() error {
	for _, c := range cs {
		if err := c.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// NewSwapClaim returns a new SwapClaim
func NewSwapClaim(owner sdk.AccAddress, rewards sdk.Coins, rewardIndexes MultiRewardIndexes) SwapClaim {
	return SwapClaim{
		BaseMultiClaim: BaseMultiClaim{
			Owner:  owner,
			Reward: rewards,
		},
		RewardIndexes: rewardIndexes,
	}
}

// GetType returns the claim's type
func (c SwapClaim) GetType() string { return SwapClaimType }

// GetReward returns the claim's reward coin
func (c SwapClaim) GetReward() sdk.Coins { return c.Reward }

// GetOwner returns the claim's owner
func (c SwapClaim) GetOwner() sdk.AccAddress { return c.Owner }

// Validate performs a basic check of a SwapClaim fields
func (c SwapClaim) Validate() error {
	if err := c.RewardIndexes.Validate(); err != nil {
		return err
	}
	return c.BaseMultiClaim.Validate()
}

// HasRewardIndex check if a claim has a reward index for the input pool ID.
func (c SwapClaim) HasRewardIndex(poolID string) (int64, bool) {
	for index, ri := range c.RewardIndexes {
		if ri.CollateralType == poolID {
			return int64(index), true
		}
	}
	return 0, false
}

// SwapClaims slice of SwapClaim
type SwapClaims []SwapClaim

// Validate checks if all the claims are valid.
func (cs SwapClaims) Validate() error {
	for _, c := range cs {
		if err := c.Validate(); err != nil {
			return err
		}
	}

	return nil
}

//// GetType returns the claim's type
//func (c SPVClaim) GetType() string { return SPVClaimType }
//
//// GetReward returns the claim's reward coin
//func (c SPVClaim) GetReward() sdk.Coins { return c.Reward }
//
//// GetOwner returns the claim's owner
//func (c SPVClaim) GetOwner() sdk.AccAddress { return c.Owner }
//
//// Validate performs a basic check of a SwapClaim fields
//func (c SPVClaim) Validate() error {
//	if err := c.RewardIndexes.Validate(); err != nil {
//		return err
//	}
//	return c.BaseMultiClaim.Validate()
//}
//
//// HasRewardIndex check if a claim has a reward index for the input pool ID.
//func (c SPVClaim) HasRewardIndex(poolID string) (int64, bool) {
//	for index, ri := range c.RewardIndexes {
//		if ri.CollateralType == poolID {
//			return int64(index), true
//		}
//	}
//	return 0, false
//}

// // SPVClaims slice of SPVClaim
type SPVClaims []SPVClaim

//
//// Validate checks if all the claims are valid.
//func (cs SPVClaims) Validate() error {
//	for _, c := range cs {
//		if err := c.Validate(); err != nil {
//			return err
//		}
//	}
//	return nil
//}

// ---------------------- Reward indexes are used internally in the store ----------------------

// NewRewardIndex returns a new RewardIndex
func NewRewardIndex(collateralType string, factor sdkmath.LegacyDec) RewardIndex {
	return RewardIndex{
		CollateralType: collateralType,
		RewardFactor:   factor,
	}
}

// Validate validates reward index
func (ri RewardIndex) Validate() error {
	if ri.RewardFactor.IsNegative() {
		return fmt.Errorf("reward factor value should be positive, is %s for %s", ri.RewardFactor, ri.CollateralType)
	}
	if strings.TrimSpace(ri.CollateralType) == "" {
		return fmt.Errorf("collateral type should not be empty")
	}
	return nil
}

// RewardIndexes slice of RewardIndex
type RewardIndexes []RewardIndex

// GetRewardIndex fetches a RewardIndex by its denom
func (ris RewardIndexes) GetRewardIndex(denom string) (RewardIndex, bool) {
	for _, ri := range ris {
		if ri.CollateralType == denom {
			return ri, true
		}
	}
	return RewardIndex{}, false
}

// Get fetches a RewardFactor by it's denom
func (ris RewardIndexes) Get(denom string) (sdkmath.LegacyDec, bool) {
	for _, ri := range ris {
		if ri.CollateralType == denom {
			return ri.RewardFactor, true
		}
	}
	return sdkmath.LegacyDec{}, false
}

// With returns a copy of the indexes with a new reward factor added
func (ris RewardIndexes) With(denom string, factor sdkmath.LegacyDec) RewardIndexes {
	newIndexes := ris.copy()

	for i, ri := range newIndexes {
		if ri.CollateralType == denom {
			newIndexes[i].RewardFactor = factor
			return newIndexes
		}
	}
	return append(newIndexes, NewRewardIndex(denom, factor))
}

// GetFactorIndex gets the index of a specific reward index inside the array by its index
func (ris RewardIndexes) GetFactorIndex(denom string) (int, bool) {
	for i, ri := range ris {
		if ri.CollateralType == denom {
			return i, true
		}
	}
	return -1, false
}

// Validate validation for reward indexes
func (ris RewardIndexes) Validate() error {
	for _, ri := range ris {
		if err := ri.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Mul returns a copy of RewardIndexes with all factors multiplied by a single value.
func (ris RewardIndexes) Mul(multiplier sdkmath.LegacyDec) RewardIndexes {
	newIndexes := ris.copy()

	for i := range newIndexes {
		newIndexes[i].RewardFactor = newIndexes[i].RewardFactor.Mul(multiplier)
	}
	return newIndexes
}

// Quo returns a copy of RewardIndexes with all factors divided by a single value.
// It uses sdkmath.LegacyDec.Quo for the division.
func (ris RewardIndexes) Quo(divisor sdkmath.LegacyDec) RewardIndexes {
	newIndexes := ris.copy()

	for i := range newIndexes {
		newIndexes[i].RewardFactor = newIndexes[i].RewardFactor.Quo(divisor)
	}
	return newIndexes
}

// Add combines two reward indexes by adding together factors with the same CollateralType.
// Any CollateralTypes unique to either reward indexes are included in the output as is.
func (ris RewardIndexes) Add(addend RewardIndexes) RewardIndexes {
	newIndexes := ris.copy()

	for _, addRi := range addend {
		found := false
		for i, origRi := range newIndexes {
			if origRi.CollateralType == addRi.CollateralType {
				found = true
				newIndexes[i].RewardFactor = newIndexes[i].RewardFactor.Add(addRi.RewardFactor)
			}
		}
		if !found {
			newIndexes = append(newIndexes, addRi)
		}
	}
	return newIndexes
}

// copy returns a copy of the reward indexes slice and underlying array
func (ris RewardIndexes) copy() RewardIndexes {
	if ris == nil { // return nil rather than empty slice when ris is nil
		return nil
	}
	newIndexes := make(RewardIndexes, len(ris))
	copy(newIndexes, ris)
	return newIndexes
}

// NewMultiRewardIndex returns a new MultiRewardIndex
func NewMultiRewardIndex(collateralType string, indexes RewardIndexes) MultiRewardIndex {
	return MultiRewardIndex{
		CollateralType: collateralType,
		RewardIndexes:  indexes,
	}
}

// GetFactorIndex gets the index of a specific reward index inside the array by its index
func (mri MultiRewardIndex) GetFactorIndex(denom string) (int, bool) {
	for i, ri := range mri.RewardIndexes {
		if ri.CollateralType == denom {
			return i, true
		}
	}
	return -1, false
}

// Validate validates multi-reward index
func (mri MultiRewardIndex) Validate() error {
	for _, rf := range mri.RewardIndexes {
		if rf.RewardFactor.IsNegative() {
			return fmt.Errorf("reward index's factor value cannot be negative: %s", rf)
		}
	}
	if strings.TrimSpace(mri.CollateralType) == "" {
		return fmt.Errorf("collateral type should not be empty")
	}
	return nil
}

// MultiRewardIndexes slice of MultiRewardIndex
type MultiRewardIndexes []MultiRewardIndex

// GetRewardIndex fetches a RewardIndex from a MultiRewardIndex by its denom
func (mris MultiRewardIndexes) GetRewardIndex(denom string) (MultiRewardIndex, bool) {
	for _, ri := range mris {
		if ri.CollateralType == denom {
			return ri, true
		}
	}
	return MultiRewardIndex{}, false
}

// Get fetches a RewardIndexes by it's denom
func (mris MultiRewardIndexes) Get(denom string) (RewardIndexes, bool) {
	for _, mri := range mris {
		if mri.CollateralType == denom {
			return mri.RewardIndexes, true
		}
	}
	return nil, false
}

// GetRewardIndexIndex fetches a specific reward index inside the array by its denom
func (mris MultiRewardIndexes) GetRewardIndexIndex(denom string) (int, bool) {
	for i, ri := range mris {
		if ri.CollateralType == denom {
			return i, true
		}
	}
	return -1, false
}

// With returns a copy of the indexes with a new RewardIndexes added
func (mris MultiRewardIndexes) With(denom string, indexes RewardIndexes) MultiRewardIndexes {
	newIndexes := mris.copy()

	for i, mri := range newIndexes {
		if mri.CollateralType == denom {
			newIndexes[i].RewardIndexes = indexes
			return newIndexes
		}
	}
	return append(newIndexes, NewMultiRewardIndex(denom, indexes))
}

// GetCollateralTypes returns a slice of containing all collateral types
func (mris MultiRewardIndexes) GetCollateralTypes() []string {
	var collateralTypes []string
	for _, ri := range mris {
		collateralTypes = append(collateralTypes, ri.CollateralType)
	}
	return collateralTypes
}

// RemoveRewardIndex removes a denom's reward interest factor value
func (mris MultiRewardIndexes) RemoveRewardIndex(denom string) MultiRewardIndexes {
	for i, ri := range mris {
		if ri.CollateralType == denom {
			// copy the slice and underlying array to avoid altering the original
			cp := mris.copy()
			return append(cp[:i], cp[i+1:]...)
		}
	}
	return mris
}

// Validate validation for reward indexes
func (mris MultiRewardIndexes) Validate() error {
	for _, mri := range mris {
		if err := mri.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// copy returns a copy of the slice and underlying array
func (mris MultiRewardIndexes) copy() MultiRewardIndexes {
	newIndexes := make(MultiRewardIndexes, len(mris))
	copy(newIndexes, mris)
	return newIndexes
}
