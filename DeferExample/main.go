package main

import "fmt"

func main() {
	fmt.Println("start")

	for i := 0; i < 10; i++ {
		defer fmt.Println("counter ", i)
	}
	fmt.Println("end")
}
