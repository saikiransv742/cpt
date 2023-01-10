package main

import (
	"fmt"
)

type user struct {
	Firstname        string
	lastname         string
	Favorite_country []string
}
type vehicle struct {
	doors string
	color string
}
type truck struct {
	a         vehicle
	fourwheel bool
}
type sedan struct {
	a      vehicle
	luxury bool
}
type squar struct {
	length int
	breath int
}
type circle struct {
	radius float64
}

func (x squar) area() int {
	return x.length * x.breath
}
func (x circle) area() float64 {
	return 3.14 * x.radius * x.radius
}

type shape interface {
	area() float64
}

func info(i shape) float64 {
	// i.area()
	return i.area()
}

func main() {
	saikiran := user{"sai", "kiran", []string{"india", "srilanka", "pakistan"}}
	fmt.Printf("%v\n", saikiran)
	fmt.Println(saikiran.Firstname)
	for i := range saikiran.Favorite_country {
		fmt.Println(i, saikiran.Favorite_country[i])
	}
	johnsnow := user{"john", "snow", []string{"usa", "austrila", "italy"}}
	fmt.Printf("%v\n", johnsnow)
	fmt.Println(johnsnow.Firstname)
	for i := range johnsnow.Favorite_country {
		fmt.Println(i, johnsnow.Favorite_country[i])
	}
	tyrionlannister := user{"tyrion", "lannister", []string{"austria", "turkey", "dubai"}}
	fmt.Printf("%v\n", tyrionlannister)
	fmt.Println(tyrionlannister.Firstname)
	for i := range tyrionlannister.Favorite_country {
		fmt.Println(i, tyrionlannister.Favorite_country[i])

		//2
		person := make(map[string]user)
		person[saikiran.lastname] = saikiran
		fmt.Printf("%v\n", person)
	}

	toyota := truck{
		fourwheel: true,
		a:         vehicle{"4door", "black"},
	}
	bmw := sedan{
		luxury: true,
		a:      vehicle{"2door", "blue"},
	}
	fmt.Println(toyota, bmw)
	fmt.Println(toyota.fourwheel)
	fmt.Println(bmw.luxury)

	w := squar{
		length: 10,
		breath: 10,
	}
	fmt.Println(w.area())
	c := circle{
		radius: 15,
	}
	fmt.Println(c.area())
}
