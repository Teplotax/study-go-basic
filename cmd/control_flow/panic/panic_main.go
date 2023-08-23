package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("start")
	panicker()
	fmt.Println("end")

	//deferPanic()
	//
	//err := forceError()
	//if err != nil {
	//	panic(err.Error())
	//}

}

func forceError() error {
	return fmt.Errorf("devision by zero")
}

func deferPanic() {

	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened") // executes deferred code before exiting
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("done panicking")
}
