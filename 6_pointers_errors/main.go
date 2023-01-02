package main

import (
	"fmt"
	"errors"
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

// struct pointers
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// go allows implicit dereference
	// instead of using (*w).balance, which is totally valid
	// we use "w",the pointer, directly:
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance{
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
