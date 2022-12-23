package main

import (
	"2calc/app"
	"fmt"
)

func main() {

	var a, b int
	fmt.Println("enter a value")
	fmt.Scanln(&a)
	fmt.Println("nter b value")
	fmt.Scanln(&b)
	fmt.Println(app.Sum(&a, &b))
	fmt.Println(app.Difference(&a, &b))
	fmt.Println(app.Multiplication(&a, &b))
	fmt.Println(app.Division(&a, &b))
}
