package main

import (
	"fmt"
	"github.com/novabankapp/wallet.data/domain"
	"time"
)

func main() {
	wa := domain.WalletLink{

		"1234",
		"1222",
		"l@m.com",
		time.Now(),
	}
	fmt.Println(wa.IsEmailLink())
}
