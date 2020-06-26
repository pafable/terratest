package main

import (
	"fmt"
	"math/rand"
)

func foo() {
	fmt.Println("A number from 1-100:", rand.Intn(100))
}

func main() {
	fmt.Println("yo")
	foo()
}
