package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}
	
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	
	return balance, nil
}



func writeBalanceToFile(balance float64) {
	balanceText := strconv.FormatFloat(balance, 'f', 2, 64)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance is", accountBalance)

		case 2:
			fmt.Println("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New Amount:", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			fmt.Println("Withdraw Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid Amount. Cannot withdraw more than your balance")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New Amount:", accountBalance)
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
		fmt.Println("Thanks for choosing our bank!")
	}
}
