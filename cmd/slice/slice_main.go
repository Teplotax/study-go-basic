package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice from the 4th element to end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice the 4th, 5th, and 6th elements
	a[5] = 42
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// Using make
	x := make([]int, 3, 100)
	fmt.Printf("Length: %v\n", len(x))
	fmt.Printf("Capacity: %v\n", cap(x))

	// Using append to add elements
	y := []int{}
	fmt.Println(y)
	fmt.Printf("Length: %v\n", len(y))
	fmt.Printf("Capacity: %v\n", cap(y))
	y = append(y, 1)
	fmt.Println(y)
	fmt.Printf("Length: %v\n", len(y))
	fmt.Printf("Capacity: %v\n", cap(y))
	y = append(y, 2, 3, 4)
	fmt.Println(y)
	fmt.Printf("Length: %v\n", len(y))
	fmt.Printf("Capacity: %v\n", cap(y))
	// You can't append a slice to a slice, but you can use a spread operator "..."
	y = append(y, []int{5, 6, 7}...)
	fmt.Println(y)
	fmt.Printf("Length: %v\n", len(y))
	fmt.Printf("Capacity: %v\n", cap(y))

	f := []int{1, 2, 3, 4, 5}
	// Removing an element from the start
	g := f[1:]
	fmt.Println(g)

	// Removing an element from the end
	h := f[:len(f)-1]
	fmt.Println(h)

	// Removing an element from the middle
	i := append(f[:2], f[3:]...)
	fmt.Println(i)

}
