package service

import (
	"account-manager-api/internal/models"
	"errors"
	"strconv"
)

type BankService struct {
	accounts map[int]*models.CheckingAccount
}

func NewBankService() *BankService {
	return &BankService{
		accounts: make(map[int]*models.CheckingAccount),
	}
}

func (b *BankService) CreateAccount(accountNumber int, owner string) {
	newAccount := &models.CheckingAccount{
		AccountNumber: accountNumber,
		Owner:         owner,
		Balance:       0.0,
	}
	b.accounts[accountNumber] = newAccount
}

func (b *BankService) DeleteAccount(accountNumber int) {
	delete(b.accounts, accountNumber)
}

func (b *BankService) Transfer(amount float64, from int, to int) error {
	accountFrom, existFrom := b.accounts[from]
	accountTo, existTo := b.accounts[to]

	if !existFrom || !existTo {
		return errors.New("uma das contas não existe")
	}

	if err := accountFrom.Withdraw(amount); err != nil {
		return err
	}

	accountTo.Deposit(amount)

	return nil
}

func (b *BankService) GetAccounts() []models.CheckingAccount {
	var accounts []models.CheckingAccount
	for _, account := range b.accounts {
		accounts = append(accounts, *account)
	}
	return accounts
}

func (b *BankService) Deposit(accountNumber int, amount float64) error {
	account, exists := b.accounts[accountNumber]
	if !exists {
		return errors.New("conta com número " + strconv.Itoa(accountNumber) + " não existe")
	}
	return account.Deposit(amount)
}

func (b *BankService) Withdraw(accountNumber int, amount float64) error {
	account, exists := b.accounts[accountNumber]
	if !exists {
		return errors.New("conta com número " + strconv.Itoa(accountNumber) + " não existe")
	}
	return account.Withdraw(amount)
}
