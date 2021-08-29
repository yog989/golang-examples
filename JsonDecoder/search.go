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

	fmt.Print("Please enter employee Id : ")
	var empId int

	fmt.Scan(&empId)

	dec := json.NewDecoder(file)

	var m []EmpInfo
	if err := dec.Decode(&m); err != nil {
		log.Fatal(err)
	}

	isFound := false
	for _, e := range m {
		if e.Id == empId {
			fmt.Printf("Id : %d, Name : %s, Location %v\n", e.Id, e.Name, e.Location)
			isFound = true
			break
		}
	}

	if !isFound {
		fmt.Printf("Employee %d not found\n", empId)
	}
}
