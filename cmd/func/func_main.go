package main

import (
	"fmt"
)

func main() {

	// Dealing with errors without panic
	d, err := divide(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	d, err = divide(5.0, 0.0)
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
