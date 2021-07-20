package main

import "fmt"

func main() {
	fmt.Print("Please enter number : ")
	var number int

	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("%d is Even number", number)
	} else {
		fmt.Printf("%d is Odd number", number)
	}
}
