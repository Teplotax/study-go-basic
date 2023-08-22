package main

import "fmt"

func main() {
	n := 3

	singleSwitch(n)
	multipleSwitch(n)
	taglessSwitch(n)
	typeSwitch(float64(n))
}

func singleSwitch(n int) {
	// You can use initializers
	switch m := n + 0; m {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Something else")
	}
}

func multipleSwitch(n int) {
	switch n {
	case 1, 2, 3:
		fmt.Println("One, Two or Three")
	case 4, 5, 6:
		fmt.Println("Four, Five or Six")
	case 7, 8, 9:
		fmt.Println("Seven, Eight or Nine")
	default:
		fmt.Println("Something else")
	}
}

func taglessSwitch(n int) {
	switch {
	case n <= 10:
		fmt.Println("less than or equal to ten")
		// The break statement is implicit, if you want to fallthrough use the keyword "fallthrough"
		fallthrough
	case n <= 20:
		fmt.Println("less than or equal to twenty")
		// You can explicitly break using the keyword "break"
		if someValidation() {
			break
		}
		fmt.Println("only prints if the validation is false")

	default:
		fmt.Println("greater than twenty")
	}
}

func typeSwitch(n interface{}) {
	switch n.(type) {
	case int:
		fmt.Println("n is an int")
	case float64:
		fmt.Println("n is a float64")
	case string:
		fmt.Println("n is a string")
	default:
		fmt.Println("n is another type")
	}
}

func someValidation() bool {
	return true
}
