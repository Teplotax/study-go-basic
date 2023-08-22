package main

import "fmt"

func main() {

	// Using pointers
	name := "Thiago Farias Reis"

	fmt.Printf("Before: %v\n", name)
	changeWithPointer(&name)
	fmt.Printf("After: %v\n", name)
	fmt.Println()
	fmt.Printf("Before: %v\n", name)
	changeWithoutPointer(name)
	fmt.Printf("After: %v\n", name)

	// Using methods
	b := greeter{
		greeting: "Sup",
		name:     "Bruno",
	}
	t := greeter{
		greeting: "Sup dude",
		name:     "Thiago",
	}

	b.greet()
	t.greet()
	b.setGreeting("Fuck yeah")
	b.greet()

}

func changeWithPointer(name *string) {
	*name = "Thiago F. Reis"
	fmt.Printf("While: %v\n", *name)
}

func changeWithoutPointer(name string) {
	name = "Thiago F. Reis"
	fmt.Printf("While: %v\n", name)
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Printf("-%v: %v!\n", g.name, g.greeting)
}

func (g *greeter) setGreeting(greeting string) {
	g.greeting = greeting
}
