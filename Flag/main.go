package main

import (
	"flag"
	"fmt"
)

func main() {
	cityName := flag.String("city", "Pune", "This is default city")
	flag.Parse()

	fmt.Println("Showing argument value passed for city (default is Pune) : ", *cityName)
}
