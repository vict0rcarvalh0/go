package main

import (
	"fibonacci/pkg/calculator"
	"fmt"
)

func main() {
	fmt.Println("Digite um número de iterações para descobrir a sua sequência de Fibonacci:")
	var n int
	fmt.Scanln(&n)
	calculator.CalculateFibonacciSequence(n)
}
