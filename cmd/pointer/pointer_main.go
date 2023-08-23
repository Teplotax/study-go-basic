package main

import "fmt"

func main() {

	//firstExample()
	pointerTypes()

}

func firstExample() {
	var a int = 42

	// This is a pointer to an integer
	var b *int

	// & is called an "address of" operator
	b = &a

	// "b" holds the location of "a"
	fmt.Println(a, b)

	// This is a dereference, it finds the location and pull the value
	fmt.Println(a, *b)

	// If we change "a", both change
	a = 27
	fmt.Println(a, *b)

	// If we change "b", both change
	*b = 14
	fmt.Println(a, *b)
}

func pointerTypes() {
	var ms *myStruct
	fmt.Println(ms) // A not initialized pointer holds the value <nil>
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
