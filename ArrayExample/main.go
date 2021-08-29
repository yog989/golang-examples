package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "First"
	name[1] = "Second"
	name[2] = "Last"

	fmt.Println(name)

	fmt.Println(name[0], name[2])

	numbers := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers)
	fmt.Println("length of array : ", len(numbers))
	fmt.Println("capacity of array : ", cap(numbers))

	var a [10]int
	fmt.Println("array a with default values as ZERO ", a)
}
