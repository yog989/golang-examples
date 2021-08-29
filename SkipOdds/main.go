package main

import "fmt"

func main() {
	numbers := [10]int{2, 3, 5, 8, 11, 13, 14, 20, 23, 50}
	fmt.Println(numbers, len(numbers), cap(numbers))

	var result []int
	fmt.Println(result, len(result), cap(result))

	// option 1
	for counter := 0; counter < len(numbers); counter++ {
		if numbers[counter]%2 == 0 {
			result = append(result, numbers[counter])
		}
	}

	fmt.Println("Result after skipping Odd numbers is : ")
	fmt.Println(result, len(result), cap(result))

	// option 2
	var option2result []int
	fmt.Println(option2result, len(option2result), cap(option2result))

	for _, number := range numbers {
		if number%2 == 0 {
			option2result = append(option2result, number)
		}
	}

	fmt.Println("Result after skipping Odd numbers is : ")
	fmt.Println(option2result, len(option2result), cap(option2result))
}
