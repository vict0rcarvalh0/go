package account

import (
	"errors"
)

type CheckingAccount struct {
	AccountNumber int
	Owner         string
	Balance       float64
}

func (c *CheckingAccount) Deposit(amount float64) error {
	c.Balance += amount
	return nil
}

func (c *CheckingAccount) Withdraw(amount float64) error {
	if amount > c.Balance {
		return errors.New("saldo insuficiente")
	}
	c.Balance -= amount
	return nil
}

func (c *CheckingAccount) CheckBalance() float64 {
	return c.Balance
}
