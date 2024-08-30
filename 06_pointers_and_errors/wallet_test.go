package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("expected %s, got %s", want, got)
		}

	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("Expected an error. No error returned")
		}

		if got != want {
			t.Errorf("expected %q, got %q", want, got.Error())
		}

	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("insufficent funds withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
