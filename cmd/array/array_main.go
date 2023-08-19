package main

import "fmt"

func main() {

	/* The three dots "..." indicate an array length that is automatically
	determined by the number of elements provided in the initializer list */

	// a is an array
	a := [...]int{1, 2, 3}
	// b is a copy of a
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := [...]int{1, 2, 3}
	// d is a reference to c
	d := &c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)

	// e is a Slice
	e := []int{1, 2, 3}
	// f is a reference to e, even without the pointer & cause e is a Slice
	f := e
	f[1] = 5
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("Length: %v\n", len(e))
	fmt.Printf("Capacity: %v\n", cap(a))

}
