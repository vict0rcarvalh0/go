package bank

import (
	account "account-manager/account"
	bank_account "account-manager/interfaces"
	"errors"
	"fmt"
)

type Bank struct {
	accounts map[int]bank_account.BankAccount
}

// Constructor function to create new bank instance
func NewBank() *Bank {
	return &Bank{
		accounts: make(map[int]bank_account.BankAccount),
	}
}

func (b *Bank) CreateAccount(accountNumber int, owner string) {
	newAccount := &account.CheckingAccount{
		AccountNumber: accountNumber,
		Owner:         owner,
		Balance:       0.0,
	}
	b.accounts[accountNumber] = newAccount
}

func (b *Bank) DeleteAccount(accountNumber int) {
	delete(b.accounts, accountNumber)
}

func (b *Bank) Transfer(amount float64, from int, to int) error {
	accountFrom, existFrom := b.accounts[from]
	accountTo, existTo := b.accounts[to]

	if !existFrom || !existTo {
		return errors.New("One of the accounts does not exist!")
	}

	if err := accountFrom.Withdraw(amount); err != nil {
		return err
	}
	accountTo.Withdraw(amount)
	return nil
}

func (b *Bank) GetAccounts() {
	for number, account := range b.accounts {
		balance := account.CheckBalance()
		fmt.Printf("Account %d - Owner: %s, Balance: %.2f\n", number, account.(account.CheckingAccount).Owner, balance)
	}
}
