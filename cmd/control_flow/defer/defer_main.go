package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	normal()

	deferMiddle()

	deferClose()
}

func normal() {
	fmt.Println("normal()")

	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")

	fmt.Println()
}

func deferMiddle() {
	fmt.Println("deferMiddle()")

	fmt.Println("start")
	defer fmt.Println()
	defer fmt.Println("middle")
	fmt.Println("end")
}

func deferClose() {

	// Defer let you open and close a resource in a single block
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// but don't use defer in loops, may cause problems

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)
}
