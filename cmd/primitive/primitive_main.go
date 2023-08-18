package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	// Booleans
	var n bool = true
	fmt.Printf("n: %v %T\n", n, n)

	// Booleans zero value
	var m bool
	fmt.Printf("m: %v %T\n", m, m)

	// Integer types
	var h int = math.MaxInt
	var i int8 = math.MaxInt8
	var j int16 = math.MaxInt16
	var k int32 = math.MaxInt32
	var l int64 = math.MaxInt64

	fmt.Printf("h: %v %T\n", h, h)
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("j: %v %T\n", j, j)
	fmt.Printf("k: %v %T\n", k, k)
	fmt.Printf("l: %v %T\n", l, l)

	// Bitwise operators
	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // 0010 AND
	fmt.Println(a | b)  // 1011 OR
	fmt.Println(a ^ b)  // 1001 XOR
	fmt.Println(a &^ b) // 0100 AND NOT

	// Bit shift
	c := 8              //2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0

	// Float types
	d := 3.14
	fmt.Printf("d: %v %T\n", d, d)
	d = 13.7e72
	fmt.Printf("d: %v %T\n", d, d)
	d = 2.1e14
	fmt.Printf("d: %v %T\n", d, d)

	// Complex numbers
	var e complex64 = 1 + 2i
	fmt.Printf("e: %v %T\n", real(e), real(e))
	fmt.Printf("e: %v %T\n", imag(e), imag(e))

	f := complex(5, 12)
	fmt.Printf("f: %v %T\n", real(f), real(f))
	fmt.Printf("f: %v %T\n", imag(f), imag(f))

	// String
	s := "this is a string"
	fmt.Printf("s: %v %T\n", s, s)
	fmt.Printf("s[2]: %v %T\n", s[2], s[2])
	fmt.Printf("s[2]: %v %T\n", string(s[2]), string(s[2]))

	// Convert string to byte slice
	t := []byte(s)
	fmt.Printf("t: %v %T\n", t, t)

	// Repeat string
	u := strings.Repeat("#", 3)
	fmt.Println(u)

}
