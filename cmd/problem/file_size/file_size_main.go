package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	filesize := 4000000000.
	fmt.Printf("%.2fGB", filesize/GB)
}
