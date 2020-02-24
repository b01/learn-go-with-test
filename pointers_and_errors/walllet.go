package bitcoin

import "fmt"

type Wallet struct {
	bal int
}

func (w *Wallet) Deposit(d int) *Wallet {
	fmt.Printf("\naddress of wallet balance in Deposit: %v\n\n", &w.bal)
	w.bal += d
	return w
}

func (w *Wallet) Balance() int {
	fmt.Printf("\naddress of wallet balance in Balance: %v\n\n", &w.bal)
	return w.bal
	//return 0
}