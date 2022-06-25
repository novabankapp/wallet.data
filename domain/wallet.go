package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type Wallet struct {
	ID               string          `json:"id"`
	UserId           string          `json:"user_id"`
	AccountId        string          `json:"account_id"`
	Balance          decimal.Decimal `json:"balance"`
	AvailableBalance decimal.Decimal `json:"available_balance"`
	CreatedAt        time.Time
}

func (w Wallet) IsNoSQLEntity() bool {
	return true
}
func (w *Wallet) GetBalance() (realBalance decimal.Decimal, availableBalance decimal.Decimal) {
	return w.Balance, w.AvailableBalance
}
