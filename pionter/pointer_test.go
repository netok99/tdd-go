package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {
	confirmBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		actual := wallet.Balance()
		if actual != expected {
			t.Errorf("actual %s, expected %s", actual, expected)
		}
	}

	confirmError := func(t *testing.T, err error, expected string) {
		t.Helper()
		if err == nil {
			t.Fatal("error dont happen.")
		}
		if err.Error() != expected {
			t.Errorf("actual %s, expected %s", err.Error(), expected)
		}
	}

	confirmOtherError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("erro inesperado recebido")
		}
	}

	t.Run("Add", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Debit(Bitcoin(10))
		confirmBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Remove", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Remove(Bitcoin(10))
		confirmBalance(t, wallet, Bitcoin(10))
		confirmOtherError(t, err)
	})

	t.Run("Remove insufficient balance", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		error := wallet.Remove(Bitcoin(100))
		confirmBalance(t, wallet, Bitcoin(20))
		confirmError(t, error, "insuficient balance")
	})
}
