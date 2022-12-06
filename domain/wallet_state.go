package domain

import (
	"github.com/gocql/gocql"
	"reflect"
)

type WalletState struct {
	IsLocked      bool       `json:"is_locked"`
	IsBlacklisted bool       `json:"is_blacklisted"`
	IsDeleted     bool       `json:"is_deleted"`
	WalletId      string     `json:"wallet_id"`
	ID            gocql.UUID `json:"id"`
}

func (w WalletState) IsNoSQLEntity() bool {
	return true
}

func (w *WalletState) CanTransact() bool {
	v := reflect.ValueOf(*w)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() == reflect.Bool {
			if field.Bool() == false {
				return false
			}
		}
	}
	return w.IsBlacklisted && w.IsLocked
}
