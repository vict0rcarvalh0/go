package main

import "fmt"

// Type Function
type transformFn func(int) int

func main() {
	// Function as Parameters
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	// Function as Values
	getTransformerFunction(&numbers)

	// Anonymous Function
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)

	// Using Closure 
	quadruple := createTransformer(4)
	quadrupled := transformNumbers(&numbers, quadruple)
	fmt.Println(quadrupled)

	// Non Recursive x Recursive Function
	fact := factorial(3)
	fmt.Println(fact)

	factRecursive := factorialRecursive(4)
	fmt.Println(factRecursive)

	// Variadic Functions
	sumResult := sum(1, 2, 3, 4, 5)
	fmt.Println(sumResult)

	splitSlices := sum(numbers...)
	fmt.Println(splitSlices)
}

// Function as parameter
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// Common functions
func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// Function as return
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

// Closure
func createTransformer(factor int) func(int) int {
	// Anonymous function
	return func(number int) int {
		return number * factor
	}
}

// Without Recursion
func factorial(number int) int {
	result := 1

	for i := 1;  i <= number; i++ {
		result *= i
	}

	return result
}

// With Recursion(function calling itself)
func factorialRecursive(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

// Variadic Functions
func sum(numbers ...int) int {
	result := 0

	for _, val := range numbers {
		result += val
	}

	return result
}