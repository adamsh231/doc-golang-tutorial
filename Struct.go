package main

import "fmt"

/**
Struct like a class in OOP
Struct is type data with plenty of field -> Can't create object with struct
Struct can inject function
*/

type CustomerPremium struct { // Struct name -> Usually using UpperCase
	Name, Address string // Field name -> Usually using UpperCase
	Age           int
	Weight        float32
	Height        float32
}

// Ordinary function
func ordFunc(customer CustomerPremium, name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

// Struct function -> injected
func (customer CustomerPremium) strctFunc(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	// One by one //
	var raycall CustomerPremium
	raycall.Name = "Raycall Arditama"
	raycall.Age = 23
	fmt.Println(raycall)

	// All -> Literal //
	// 1 . First  -> use this instead of second step
	var nana = CustomerPremium{
		Name: "Roudhlotul Jannah",
		Age:  22,
	}
	fmt.Println(nana)

	// 2. Second
	var faiz = CustomerPremium{"Faiz Adhitya", "Malang", 22, 0, 0}
	fmt.Println(faiz)

	// Struct Method //
	// 1. Ordinary function
	ordFunc(nana, "Faiz")

	// 2. Struct function
	nana.strctFunc("Faiz")
}
