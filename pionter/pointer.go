package pointer

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Debit(quantity Bitcoin) {
	w.balance += quantity
}

var insuficienteBalance = errors.New("insuficient balance")

func (w *Wallet) Remove(quantity Bitcoin) error {
	if quantity > w.balance {
		return insuficienteBalance
	}
	w.balance -= quantity
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
