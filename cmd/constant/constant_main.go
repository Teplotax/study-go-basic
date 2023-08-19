package main

import "fmt"

const (
	_ = iota + 5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	const myConst int = 42
	fmt.Printf("myConst: %v %T\n", myConst, myConst)

	var specialistType = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
}
