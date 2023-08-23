package main

import "fmt"

func main() {
	//basic()
	//incrementerInterface()
	//embeddedInterfaces()
	typeSwitch()
}

// ===================================================================================
func basic() {

	// "w" is polymorphic, we can specify any type tha implements the methods of the interface Writer
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

// Writer The interface, this is the naming convention for single method interfaces
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// The implementation
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// ===================================================================================
func incrementerInterface() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

// IntCounter : This is a type alias for int
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// ===================================================================================
func embeddedInterfaces() {
	myInt := IntCounter(0)
	var inc IncrementerDecrementer = &myInt
	for i := 0; i < 3; i++ {
		fmt.Println(inc.Increment())
	}
	for i := 0; i < 3; i++ {
		fmt.Println(inc.Decrement())
	}

}

type Decrementer interface {
	Decrement() int
}

func (ic *IntCounter) Decrement() int {
	*ic--
	return int(*ic)
}

type IncrementerDecrementer interface {
	Incrementer
	Decrementer
}

// ===================================================================================
func typeSwitch() {
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}
}

//===================================================================================
