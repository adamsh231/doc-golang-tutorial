package main

import "fmt"

func main() {
	var name = "Adam"

	if name == "Adam" {
		fmt.Println("Correct")
	} else if name == "Unknown" {
		fmt.Println("Hi, How do you do?")
	} else {
		fmt.Println("Hi, Nice to meet you?")
	}

	// Short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama Kepanjangan")
	} else {
		fmt.Println("Nama Kependekan")
	}
}
