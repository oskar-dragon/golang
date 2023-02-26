package pointersanderrors

import (
	"testing"
)


func TestWallet(t *testing.T) {
	t.Run("returns current balance", func(t *testing.T) {
		wallet := Wallet{
			balance: 10,
		}
		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, got, want)
	})

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{
			balance: 10,
		}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(20)

		assertBalance(t, got, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: 10,
		}
		err := wallet.Withdraw(Bitcoin(5))
		got := wallet.Balance()
		want := Bitcoin(5)

		assertNoError(t, err)
		assertBalance(t, got, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet {
			balance: 10,
		}
		err := wallet.Withdraw(Bitcoin(20))

		asserError(t, err, ErrInsufficientFunds.Error())
	})

}