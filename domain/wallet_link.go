package domain

import (
	"github.com/gocql/gocql"
	"net/mail"
	"time"
)

type WalletLink struct {
	WalletId string     `json:"wallet_id"`
	ID       gocql.UUID `json:"id"`
	Value    string     `json:"value"`
	LinkDate time.Time
}

func (w *WalletLink) IsEmailLink() bool {
	_, err := mail.ParseAddress(w.Value)
	if err != nil {
		return false
	}
	return true
}
