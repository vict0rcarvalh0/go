package calculator

import (
	"fmt"
)

func CalculateFibonacciSequence(n int) []int {
	var slice []int

	for i := 0; i < n; i++ {
		if i < 2 {
			slice = append(slice, i)
		} else {
			slice = append(slice, slice[i-1]+slice[i-2])
		}
	}

	fmt.Println("A sequencia de Fibonacci para", n, "iterações é:")
	fmt.Println(slice)
	return slice
}
