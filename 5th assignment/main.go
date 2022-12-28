package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	// 2nd problem
	i := 1
	for i <= 50 {
		if i%2 != 0 {
			fmt.Println(i)
		}
		i++
	}
	// 3rd problem
	j := 1
	for {
		if j <= 50 {
			if j%2 == 0 {
				fmt.Println(j)
			}
			j++
		}
		if j == 50 {
			break
		}
	}

	// 4th problem
	for i := 50; i <= 105; i++ {
		a := i / 6
		fmt.Println("quotient is ", a)

	}
	// 5th problem
	var b string
	fmt.Println(" enter the input")
	fmt.Scan(&b)
	if b == "golangtutorial" {
		fmt.Println("welcome")
	} else {
		fmt.Println("end")
	}
	// 6th problem
	for i := 1; i <= 80; i++ {
		if i%2 == 0 && i%4 == 0 {
			fmt.Println("golang tutorial")
		} else if i%2 == 0 {
			fmt.Println("golang")
		} else if i%4 == 0 {
			fmt.Println("tutorial")
		} else if i%2 != 0 && i%4 != 0 {
			fmt.Println(i)

		}

	}

}
