package main

import "fmt"

func main() {
	fmt.Println("Printing 10 Odd numbers")

	// Option 1
	for number := 1; number < 20; number += 2 {
		fmt.Println("Odd number ", number)
	}

	// Option 2
	fmt.Println("-------------------")

	odd := 0
	counter := 0
	for {
		odd++
		if odd%2 == 0 {
			continue
		}

		if counter == 10 {
			break
		}
		counter++
		fmt.Println("Odd number", odd)
	}
}
