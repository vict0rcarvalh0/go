package main

import "fmt"

func main() {
	resultDynamicOperation := addDynamic(1, 2)
	fmt.Println(resultDynamicOperation)
	//result + 1 // error: invalid operation: result + 1 (mismatched types interface {} and int)

	resultGenericOperation := addGenerics(1, 2)
	fmt.Println(resultGenericOperation)
	fmt.Println(resultGenericOperation + 1)
}

// using dynamic types
func addDynamic(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	return nil
}

// using generics
func addGenerics[T int|float64|string](a, b T) T {
	return a + b
}