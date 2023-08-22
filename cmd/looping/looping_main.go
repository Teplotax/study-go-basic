package main

import "fmt"

func main() {
	twoVarLoop()
	while()
	breakTag()
	collection()

}

func normalLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func twoVarLoop() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Printf("i: %v , j: %v\n", i, j)
	}
}

func while() {
	i := 5
	boolean := true
	for boolean {
		if i > 0 {
			fmt.Println("Not yet")
			i--
		} else {
			fmt.Println("Now it's over")
			boolean = false
		}
	}
}

func breakTag() {
Loop:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 8 {
				break Loop
			}
		}
	}
}

func collection() {
	s := []int{1, 2, 3}
	for key, value := range s {
		fmt.Printf("index[%v] = %v\n", key, value)
	}

	// If you don't care about the key
	for _, value := range s {
		fmt.Printf("value = %v\n", value)
	}
}
