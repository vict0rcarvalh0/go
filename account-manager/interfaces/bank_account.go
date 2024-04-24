package interfaces

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	CheckBalance() float64
}
