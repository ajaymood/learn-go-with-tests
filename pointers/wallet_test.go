package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit and check balance", func(t *testing.T) {
		wallet := Wallet{}
		b := Bitcoin(10)
		wallet.Deposit(b)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw and check balance", func(t *testing.T) {
		b := Bitcoin(10)
		wallet := Wallet{balance: b}
		err := wallet.Withdraw(b)
		if err != nil {
			t.Errorf("Error during withdraw: %s", err.Error())
		}
		want := Bitcoin(0)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {

		wallet := Wallet{}
		err := wallet.Withdraw(Bitcoin(10))
		wantBalance := Bitcoin(0)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, wantBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("error expected but didn't get any")
	}
	if got.Error() != want.Error() {
		t.Errorf("got: %s want: %s", got.Error(), want.Error())
	}
}
