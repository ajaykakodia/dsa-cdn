package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are in searching folder")
	var input string
	for {
		fmt.Print("Please provide your input: ")
		fmt.Scanln(&input)
		if input == "-1" {
			break
		}
		fmt.Println("Entered Input is : ", input)
	}
}
