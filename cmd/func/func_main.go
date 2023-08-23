package main

import (
	"fmt"
)

func main() {

	//noPanic()
	//doPanic()

	// Accept any amount of int values
	result := variadicArgument("Some text here", 1, 2, 3, 4, 5)
	fmt.Printf("Printing the result again: %v\n", result)

	// Anonimous function
	func() {
		fmt.Println("This is an anonimous function")
	}()

	// Function as a variable
	f := func() {
		fmt.Println("Invokes with f()")
	}
	fmt.Println("Do something else")
	f()
}

func noPanic() {
	// Dealing with errors without panic
	d, err := divide(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func doPanic() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// The variadic argument must be the last one
func variadicArgument(text string, values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(text)
	fmt.Printf("Sum of %v = %v\n", values, result)
	return result
}
