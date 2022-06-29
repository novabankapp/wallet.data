package models

type WalletProjection struct {
	ID                 string `json:"id"`
	WalletID           string `json:"wallet_id,omitempty"`
	UserID             string `json:"wallet_id,omitempty"`
	Wallet             string `json:"wallet"`
	WalletState        string `json:"wallet_state"`
	WalletTransactions string `json:"wallet_transactions"`
}

func (w WalletProjection) IsNoSQLEntity() bool {
	return true
}
