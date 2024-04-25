package main

import (
	"account-manager/bank"
	"fmt"
)

func main() {
	bank := bank.NewBank()

	for {
		fmt.Println("\nEscolha uma operação:")
		fmt.Println("1. Criar conta")
		fmt.Println("2. Deletar conta")
		fmt.Println("3. Depositar em uma conta")
		fmt.Println("4. Sacar de uma conta")
		fmt.Println("5. Transferir entre contas")
		fmt.Println("6. Exibir contas")
		fmt.Println("0. Sair")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Create
			var accountNumber int
			var owner string
			fmt.Println("Digite o número da conta:")
			fmt.Scanln(&accountNumber)
			fmt.Println("Digite o nome do titular:")
			fmt.Scanln(&owner)
			bank.CreateAccount(accountNumber, owner)
			fmt.Println("Conta criada com sucesso.")

		case 2:
			// Delete
			var accountNumber int
			fmt.Println("Digite o número da conta a ser deletada:")
			fmt.Scanln(&accountNumber)
			bank.DeleteAccount(accountNumber)
			fmt.Println("Conta deletada com sucesso.")

		case 3:
			// Deposit
			var accountNumber int
			var amount float64
			fmt.Println("Digite o número da conta:")
			fmt.Scanln(&accountNumber)
			fmt.Println("Digite o valor a ser depositado:")
			fmt.Scanln(&amount)
			err := bank.Deposit(accountNumber, amount)
			if err != nil {
				fmt.Println("Erro ao depositar:", err)
			} else {
				fmt.Println("Depósito realizado com sucesso.")
			}

		case 4:
			// Withdraw
			var accountNumber int
			var amount float64
			fmt.Println("Digite o número da conta:")
			fmt.Scanln(&accountNumber)
			fmt.Println("Digite o valor a ser sacado:")
			fmt.Scanln(&amount)
			err := bank.Withdraw(accountNumber, amount)
			if err != nil {
				fmt.Println("Erro ao sacar:", err)
			} else {
				fmt.Println("Saque realizado com sucesso.")
			}

		case 5:
			// Transfer
			var fromAccount, toAccount int
			var amount float64
			fmt.Println("Digite o número da conta de origem:")
			fmt.Scanln(&fromAccount)
			fmt.Println("Digite o número da conta de destino:")
			fmt.Scanln(&toAccount)
			fmt.Println("Digite o valor a ser transferido:")
			fmt.Scanln(&amount)
			err := bank.Transfer(amount, fromAccount, toAccount)
			if err != nil {
				fmt.Println("Erro ao transferir:", err)
			} else {
				fmt.Println("Transferência realizada com sucesso.")
			}

		case 6:
			// Get Accounts
			bank.GetAccounts()

		case 0:
			// Leave
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
