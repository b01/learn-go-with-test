package bitcoin

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T)  {
	wlt := Wallet{}

	wlt.Deposit(10)
	w := 10
	g := wlt.Balance()

	fmt.Printf("\naddress of wallet balance in test: %v\n\n", &wlt.bal)

	if w != g {
		t.Errorf("want %v got %v", w, g)
	}
}
