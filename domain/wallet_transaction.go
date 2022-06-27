package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type WalletTransaction struct {
	DebitWalletId  string          `json:"debit_wallet_id"`
	CreditWalletId string          `json:"credit_wallet_id"`
	Amount         decimal.Decimal `json:"amount"`
	CreatedAt      time.Time       `json:"created_at"`
	Description    string          `json:"description"`
}
