package main

import (
	"fmt"

	ASSIGNMENTS "ASSIGNMENTS/calc"
)

func main() {

	var a, b int
	fmt.Println("enter a value")
	fmt.Scanln(&a)
	fmt.Println("nter b value")
	fmt.Scanln(&b)
	fmt.Println(ASSIGNMENTS.Sum(a, b))
	fmt.Println(ASSIGNMENTS.Difference(a, b))
	fmt.Println(ASSIGNMENTS.Multiplication(a, b))
	fmt.Println(ASSIGNMENTS.Division(a, b))
}
