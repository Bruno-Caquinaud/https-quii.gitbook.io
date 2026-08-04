package pointersErrors

import (
	"fmt"
	"errors"
)


var ErrorInsuficientFunds = errors.New("Unable to procceed Withdraw operation, balance is too low")
type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w * Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w * Wallet) Balance() Bitcoin {
	return w.balance
}

func (w * Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return ErrorInsuficientFunds
	}

	w.balance -= amount
	return nil
}