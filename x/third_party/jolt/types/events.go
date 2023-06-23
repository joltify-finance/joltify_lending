package types

// Event types for jolt module
const (
	EventTypeHardDeposit          = "hard_deposit"
	EventTypeHardWithdrawal       = "hard_withdrawal"
	EventTypeJoltBorrow           = "hard_borrow"
	EventTypeHardLiquidation      = "hard_liquidation"
	EventTypeHardRepay            = "hard_repay"
	AttributeValueCategory        = ModuleName
	AttributeKeyDepositor         = "depositor"
	AttributeKeyBorrower          = "borrower"
	AttributeKeyBorrowCoins       = "borrow_coins"
	AttributeKeySender            = "sender"
	AttributeKeyRepayCoins        = "repay_coins"
	AttributeKeyLiquidatedOwner   = "liquidated_owner"
	AttributeKeyLiquidatedCoins   = "liquidated_coins"
	AttributeKeyKeeper            = "keeper"
	AttributeKeyKeeperRewardCoins = "keeper_reward_coins"
	AttributeKeyOwner             = "owner"
)
