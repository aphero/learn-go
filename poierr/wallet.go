package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(a Bitcoin) {
	w.balance += a
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(a Bitcoin) error {
	if a > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= a
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
