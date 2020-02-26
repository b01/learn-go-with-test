package bitcoin

import (
	"testing"
)

func TestBitcoin(t *testing.T)  {
	t.Run("Deposit", func(t *testing.T) {
		wlt := Wallet{}
		w := Bitcoin(10)
		wlt.Deposit(w)
		g := wlt.Balance()

		if w != g {
			t.Errorf("want %s got %s", w, g)
		}
	})
}
