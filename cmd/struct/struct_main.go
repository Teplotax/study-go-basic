package main

import (
	"fmt"
	"reflect"
)

// Doctor - Usually you'll use a struct to define a type
type Doctor struct {
	// PascalCase for export, camelCase for package access
	Number     int
	ActorName  string
	Companions []string
}

func main() {
	aDoctor := Doctor{
		Number:    3,
		ActorName: "Jon Pertwee",
		Companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor.Companions[1])

	// For short-lived vars you can assign a struct directly
	doctor := struct{ name string }{name: "John Pertwee"}
	// Unlike a slice, the "beaver operator" makes a copy
	doctor2 := doctor
	doctor2.name = "Tom Baker"
	fmt.Println(doctor)
	fmt.Println(doctor2)

	// To make a reference to the same struct use the address operator &
	doctor3 := &doctor
	doctor3.name = "Dolittle"
	fmt.Println(doctor)
	fmt.Println(doctor3)

	embedding()
	tag()

}

// Embedding is used for composition, there is no inheritance in go
func embedding() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	// You need to be aware of the embedding when using the literal syntax to instantiate a bird
	bird := Bird{
		Animal:   Animal{Name: "Penguin", Origin: "Antarctic"},
		SpeedKPH: 2,
		CanFly:   false,
	}
	fmt.Println(bird)
}

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func tag() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
