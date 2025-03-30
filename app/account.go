package app

import (
	"errors"
	"sync"
)

type Account struct {
	User      *User
	Portfolio *Portfolio
	Balance   float32
	mu        sync.RWMutex
}

func NewAccount(user *User, portfolio *Portfolio) *Account {
	return &Account{
		User:      user,
		Portfolio: portfolio,
	}
}

func (a *Account) Deposit(amount float32) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance += amount

}

func (a *Account) Withdraw(amount float32) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Balance < amount {
		return errors.New("Insufficient funds in account")
	}

	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float32 {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.Balance
}
