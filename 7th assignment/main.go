package main

import (
	"fmt"
	"os"
)

func CreatFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	defer file.Close()

	_, writeErr := file.WriteString("hello golang")
	if writeErr != nil {
		fmt.Println(err.Error)
		return
	}
	{
		data_by, err := os.ReadFile("text.txt")
		if err != nil {
			fmt.Printf("%s,err")
		}

		data_str := fmt.Sprintf("%s", data_by)
		fmt.Println(data_str)
	}

}

func main() {
	CreatFile("text.txt")
}
