package main

import "fmt"

func namedresult(passed int) (x, y int) {
	x = passed - 1
	y = passed + 1
	return
}

func main() {
	a, b := namedresult(20)
	fmt.Println("Named result for ", a, b)
}
