package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("can't withdraw more than balance")

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.balance += b
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(b Bitcoin) error {
	if w.balance-b < 0 {
		// return fmt.Errorf("can't withdraw more than balance")
		return errors.New("can't withdraw more than balance")
	}
	w.balance -= b
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
