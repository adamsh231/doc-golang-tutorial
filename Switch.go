package main

import "fmt"

func main() {
	var name = "Adam"

	switch name {
	case "Adam":
		fmt.Println("Correct")
	default:
		fmt.Println("Other")
	}

	// Short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Panjang")
	case false:
		fmt.Println("Pendek")
	}

	// Switch kondisi didalam
	var length2 = len(name)

	switch {
	case length2 > 5:
		fmt.Println("Panjang")
	case length2 < 5:
		fmt.Println("Pendek")
	}
}
