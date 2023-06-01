package common

import (
	"time"
)

const (
	WITHDRAW = iota
	PAYPRINCIPAL
)

type SPV struct {
	PoolInfo struct {
		Index           string    `json:"index"`
		PoolName        string    `json:"pool_name"`
		LinkedProject   int       `json:"linked_project"`
		OwnerAddress    string    `json:"owner_address"`
		Apy             string    `json:"apy"`
		PrincipalPaid   bool      `json:"principal_paid"`
		PayFreq         int       `json:"pay_freq"`
		ReserveFactor   string    `json:"reserve_factor"`
		PoolNFTIds      []string  `json:"pool_nFT_ids"`
		LastPaymentTime time.Time `json:"last_payment_time"`
		PoolStatus      string    `json:"pool_status"`
		BorrowedAmount  struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"borrowed_amount"`
		PoolInterest  string `json:"pool_interest"`
		ProjectLength string `json:"project_length"`
		UsableAmount  struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"usable_amount"`
		TargetAmount struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"target_amount"`
		PoolType              string `json:"pool_type"`
		EscrowInterestAmount  string `json:"escrow_interest_amount"`
		EscrowPrincipalAmount struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"escrow_principal_amount"`
		WithdrawProposalAmount struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"withdraw_proposal_amount"`
		ProjectDueTime                         time.Time     `json:"project_due_time"`
		WithdrawAccounts                       []interface{} `json:"withdraw_accounts"`
		TransferAccounts                       []interface{} `json:"transfer_accounts"`
		WithdrawRequestWindowSeconds           int           `json:"withdraw_request_window_seconds"`
		PoolLockedSeconds                      int           `json:"pool_locked_seconds"`
		PoolTotalBorrowLimit                   int           `json:"pool_total_borrow_limit"`
		CurrentPoolTotalBorrowCounter          int           `json:"current_pool_total_borrow_counter"`
		PoolCreatedTime                        time.Time     `json:"pool_created_time"`
		PoolFirstDueTime                       time.Time     `json:"pool_first_due_time"`
		GraceTime                              string        `json:"grace_time"`
		NegativeInterestCounter                int           `json:"negative_interest_counter"`
		LiquidationRatio                       string        `json:"liquidation_ratio"`
		TotalLiquidationAmount                 string        `json:"total_liquidation_amount"`
		PrincipalPaymentExchangeRatio          string        `json:"principal_payment_exchange_ratio"`
		PrincipalWithdrawalRequestPaymentRatio string        `json:"principal_withdrawal_request_payment_ratio"`
		PoolDenomPrefix                        string        `json:"pool_denom_prefix"`
		InterestPrepayment                     interface{}   `json:"interest_prepayment"`
		TransferAccountsNumber                 int           `json:"transfer_accounts_number"`
	} `json:"pool_info"`

	Depositor struct {
		InvestorID       string `json:"investor_id"`
		DepositorAddress string `json:"depositor_address"`
		PoolIndex        string `json:"pool_index"`
		LockedAmount     struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"locked_amount"`
		WithdrawalAmount struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"withdrawal_amount"`
		IncentiveAmount struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"incentive_amount"`
		LinkedNFT       []string `json:"linkedNFT"`
		DepositType     string   `json:"deposit_type"`
		PendingInterest struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"pending_interest"`
		TotalPaidLiquidationAmount string `json:"total_paid_liquidation_amount"`
	} `json:"depositor"`

	ClaimableInterestAmount struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"claimable_interest_amount"`

	Class struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Symbol      string `json:"symbol"`
		Description string `json:"description"`
		URI         string `json:"uri"`
		URIHash     string `json:"uri_hash"`
		Data        struct {
			Type          string    `json:"@type"`
			PoolIndex     string    `json:"pool_index"`
			Apy           string    `json:"apy"`
			PayFreq       int       `json:"pay_freq"`
			IssueTime     time.Time `json:"issue_time"`
			BorrowDetails []struct {
				BorrowedAmount struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"borrowed_amount"`
				TimeStamp     time.Time `json:"time_stamp"`
				ExchangeRatio string    `json:"exchange_ratio"`
			} `json:"borrow_details"`
			MonthlyRatio string `json:"monthly_ratio"`
			InterestSPY  string `json:"interest_sPY"`
			Payments     []struct {
				PaymentTime   time.Time `json:"payment_time"`
				PaymentAmount struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"payment_amount"`
				BorrowedAmount struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"borrowed_amount"`
			} `json:"payments"`
			InterestPaid struct {
				Denom  string `json:"denom"`
				Amount string `json:"amount"`
			} `json:"interestPaid"`
			AccInterest struct {
				Denom  string `json:"denom"`
				Amount string `json:"amount"`
			} `json:"acc_interest"`
			LiquidationItems           []interface{} `json:"liquidation_items"`
			TotalPaidLiquidationAmount string        `json:"total_paid_liquidation_amount"`
		} `json:"data"`
	} `json:"class"`

	Balances []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"balances"`

	Price struct {
		MarketID string `json:"market_id"`
		Price    string `json:"price"`
	} `json:"price"`
}
