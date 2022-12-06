package domain

import (
	"github.com/gocql/gocql"
	"github.com/shopspring/decimal"
	"sync"
	"time"
)

type Wallet struct {
	ID               gocql.UUID      `json:"id"`
	UserId           string          `json:"user_id"`
	AccountId        string          `json:"account_id"`
	Balance          decimal.Decimal `json:"balance"`
	AvailableBalance decimal.Decimal `json:"available_balance"`
	CreatedAt        time.Time
	Lock             sync.RWMutex
}

func (w Wallet) IsNoSQLEntity() bool {
	return true
}
func (w *Wallet) GetBalance() (realBalance decimal.Decimal, availableBalance decimal.Decimal) {
	return w.Balance, w.AvailableBalance
}
