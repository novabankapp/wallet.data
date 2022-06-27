package models

import "github.com/novabankapp/wallet.data/domain"

type WalletProjection struct {
	ID                 string                     `json:"id"`
	WalletID           string                     `json:"wallet_id,omitempty"`
	Wallet             domain.Wallet              `json:"wallet"`
	WalletState        domain.WalletState         `json:"wallet_state"`
	WalletTransactions []domain.WalletTransaction `json:"wallet_transactions"`
}

func (w WalletProjection) IsNoSQLEntity() bool {
	return true
}
