package banking

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount : creates new accoutn with 0 balance
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // return memory address to avoid
}

// declare error type
var errNoMoney = errors.New("Can't withdraw") // name of error variable is recommended to start with 'err'

// method is not function?

func (a *Account) Deposit(amount int) { // (theAccount Account -> 'a' is a receiver and should start with small letter!)
	a.balance += amount
}

// Get current balance of account a
func (a *Account) Balance() int {
	return a.balance
}

// withdraw
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount

	return nil
}

// ChangeOwner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner
func (a *Account) Owner() string {
	return a.owner
}

// String method
func (a *Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
