package main

import "fmt"

func main() {

	// Creating a map with values
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	// Adding a pair to a map
	statePopulations["Georgia"] = 10310371

	fmt.Println(statePopulations["Georgia"])

	// Deleting a pair from a map
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	// Creating an empty map
	stateBirthRate := make(map[string]int)
	fmt.Println(stateBirthRate)

	// The map doesn't throw an error switch we misspelled the key
	fmt.Println(statePopulations["Oho"]) // it returns a zero

	// The zero amy or may not be a true value, we can use the optional "ok" to check
	pop, ok := statePopulations["Oho"]
	fmt.Println(pop, ok)

	pop, ok = statePopulations["Ohio"]
	fmt.Println(pop, ok)

	// Length of a map
	fmt.Println(len(statePopulations))
}
