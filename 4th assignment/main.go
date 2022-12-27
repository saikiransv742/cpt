package main

import "fmt"

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", a)

	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%T\n", b)
	c := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(c[0:5])
	fmt.Println(c[5:])
	fmt.Println(c[2:7])
	fmt.Println(c[1:6])

	d := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}
	d = append(d, 53)
	fmt.Println("the slice", d)
	d = append(d, 53, 54, 55)
	fmt.Println("the slice", d)
	e := []int{56, 57, 58, 59, 60}
	d = append(d, e...)
	fmt.Println(d)

}
