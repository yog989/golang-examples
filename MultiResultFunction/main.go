package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("first", "second")
	fmt.Println("Swapped Result is ", a, b)
}
