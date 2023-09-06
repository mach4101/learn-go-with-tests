package pointer

import (
	"errors"
	"fmt"
)

type Bitcon int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcon
}

func (w *Wallet) Deposit(amount Bitcon) {
	w.balance += amount
}

var InsufficientFundsError = errors.New("cannot Withdraw, Insufficient Funds")

func (w *Wallet) Withdraw(amount Bitcon) error {
	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}

func (w Wallet) Balance() Bitcon {
	return w.balance
}

func (b Bitcon) String() string {
	return fmt.Sprintf("%d BTC", b)
}
