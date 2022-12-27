package main

import "fmt"

func main() {
	cities := [4]string{"Newyork", "Chicago", "Boston", "seattle"}

	fmt.Println("givin cities: ", cities)
	updateThirdElement(&cities[2])
	fmt.Println("updated cities: ", cities)

}

func updateThirdElement(cities *string) {
	*cities = "Texas"
	fmt.Println(*cities)

}
