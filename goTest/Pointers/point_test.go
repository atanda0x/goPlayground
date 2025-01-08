package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		got := wallet.Balance()

		fmt.Printf("address of balance in test is %p\n", &wallet.balance)

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()

		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}

	})
}
