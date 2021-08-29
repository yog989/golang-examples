package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type EmpInfo struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Id       int    `json:"empId"`
}

func main() {
	file, err := os.Open("employee_info.json")

	//handle error, if it fails to open file
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Println("Successfully opened employee_info.json")
	dec := json.NewDecoder(file)

	var m []EmpInfo
	if err := dec.Decode(&m); err != nil {
		log.Fatal(err)
	}

	for _, e := range m {
		fmt.Printf("%d : %s, %v\n", e.Id, e.Name, e.Location)
	}
}
