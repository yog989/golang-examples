package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // point to i

	*p = 21

	fmt.Println(i, j, *p)

	p = &j         // Point to j
	*p = *p / 37   // devide j through the pointer
	fmt.Println(j) // see the new value of j
}
