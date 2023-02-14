package types

const (
	EventTypeCreatePool        = "pool_created"
	EventTypeDeposit           = "deposit_created"
	EventTypeAddInvestors      = "add_investors"
	EventTypeBorrow            = "borrow"
	EventTypeRepayInterest     = "repay_interest"
	EventTypeClaimInterest     = "claim_interest"
	EventTypePayPrincipal      = "pay_principal"
	EventTypeWithdrawPrincipal = "withdraw_principal"
)
const (
	AttributeCreator = "creator"
	AttributeAmount  = "token"
)
