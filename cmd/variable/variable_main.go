package main

import (
	"fmt"
	"strconv"
)

// Declaring variables in the package level
var l float32 = 123

// Declaring multiple variables with a var block
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	// Assigning type and value separately.
	var i int
	i = 42
	fmt.Printf("i: %v %T\n", i, i)

	// Assigning type and value together.
	var j int = 27
	fmt.Printf("j: %v %T\n", j, j)

	// Short variable declaration.
	k := 99
	fmt.Printf("k: %v %T\n", k, k)

	// Shadowing, declaration in the innermost scope takes precedence
	fmt.Printf("l: %v %T\n", l, l)
	l := 123
	fmt.Printf("l: %v %T\n", l, l)

	// Conversion number type
	m := float32(j)
	fmt.Printf("m: %v %T\n", m, m)

	// Convert int to string using strconv
	n := strconv.Itoa(i)
	fmt.Printf("n: %v %T\n", n, n)

}
