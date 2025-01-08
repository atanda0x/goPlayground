package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		AssertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		AssertNoErr(t, err)
		AssertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw Insuffiecient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(100))

		AssertBalance(t, wallet, startBalance)
		AssertErr(t, err, ErrInsufficientFunds.Error())
	})
}

func AssertErr(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an err but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func AssertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}

func AssertNoErr(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an err but didn't want one")
	}
}
