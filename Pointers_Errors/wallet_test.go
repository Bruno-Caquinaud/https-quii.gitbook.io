package pointersErrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, got Bitcoin) {
		t.Helper()
		want := wallet.Balance()

		if got != want {
			t.Errorf("got %s != want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		if want == nil {
			t.Errorf("An error should be raised in this case")
		}

		if want != got {
			t.Errorf("Expected error does not match")
		}
	}

	assertNoError := func(t testing.TB, got error) {
		if got != nil {
			t.Errorf("Expected error no error")
		}
	}
	t.Run("Test 1 : Deposit Method", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		got := Bitcoin(10)

		assertBalance(t, wallet, got)
	})

	t.Run("Test 2 : Withdraw Method", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(15))
		got := Bitcoin(5)

		assertNoError(t, err)
		assertBalance(t, wallet, got)
	})

	t.Run("Test 3 : Withdraw Method Negative Solde", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(40))

		assertError(t, err, ErrorInsuficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}
