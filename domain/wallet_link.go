package domain

import (
	"net/mail"
	"time"
)

type WalletLink struct {
	WalletId string `json:"wallet_id"`
	ID       string `json:"id"`
	Value    string `json:"value"`
	LinkDate time.Time
}

func (w *WalletLink) IsEmailLink() bool {
	_, err := mail.ParseAddress(w.Value)
	if err != nil {
		return false
	}
	return true
}
