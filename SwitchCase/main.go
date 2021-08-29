package main

import "fmt"

func main() {
	i := true

	switch i {
	case true:
		fmt.Println("value is true ", i)
		fallthrough
	case false:
		fmt.Println("value is false ", i)
		fallthrough
	default:
		fmt.Println("value is not true/false ", i)
	}
}
