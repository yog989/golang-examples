package main

import "fmt"

func main() {
	var marks int

	fmt.Print("Please enter marks : ")
	fmt.Scan(&marks)

	// Option 1 - with if else
	if marks >= 60 {
		fmt.Printf("Result : FIRST CLASS, Marks %d", marks)
	} else if marks >= 40 && marks < 60 {
		fmt.Printf("Result : PASS, marks %d", marks)
	} else {
		fmt.Printf("Result : FAIL, marks %d", marks)
	}

	// Option 2 - with switch case
	switch {
	case marks >= 60:
		fmt.Printf("Result : FIRST CLASS , Marks %d", marks)
	case marks >= 40 && marks < 60:
		fmt.Printf("Result : PASS , Marks %d", marks)
	default:
		fmt.Printf("Result : FAIL , Marks %d", marks)
	}
}
