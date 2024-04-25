package bank

import (
	checking_account "account-manager/account"
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
	newAccount := &checking_account.CheckingAccount{
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
		return errors.New("one of the accounts does not exist")
	}

	if err := accountFrom.Withdraw(amount); err != nil {
		return err
	}

	err = accountTo.Deposit(amount)
	if err != nil {
		accountFrom.Deposit(amount)
		return err
	}

	return nil
}

func (b *Bank) GetAccounts() {
	for number, account := range b.accounts {
		balance := account.CheckBalance()
		fmt.Printf("Account %d - Owner: %s, Balance: %.2f\n", number, account.(*checking_account.CheckingAccount).Owner, balance)
	}
}

func (b *Bank) Deposit(accountNumber int, amount float64) error {
	account, exists := b.accounts[accountNumber]
	if !exists {
		return fmt.Errorf("conta com número %d não existe", accountNumber)
	}
	return account.Deposit(amount)
}

func (b *Bank) Withdraw(accountNumber int, amount float64) error {
	account, exists := b.accounts[accountNumber]
	if !exists {
		return fmt.Errorf("conta com número %d não existe", accountNumber)
	}
	return account.Withdraw(amount)
}
