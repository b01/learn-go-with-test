package bitcoin

import (
	"fmt"
)

type Wallet struct {
	bal Bitcoin
}

func (w *Wallet) Deposit(d Bitcoin) *Wallet {
	w.bal += d
	return w
}

func (w *Wallet) Balance() Bitcoin {
	return w.bal
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}