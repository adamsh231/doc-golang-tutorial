package main

import "fmt"

/**
Interface doesnt unlike other interface in OOP
Interface like other Abstract in OOP
Interface is also data type
Interface containt method declaration
Interface usually used for contract -> every function or field declared in interface is interface it self
									-> if variable or function has same name in interface's contract
									-> variable or function is interface it self (Automatic)
									-> not manually declare is interface or not like any other OOP
*/

type HasName interface {
	getName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

// Automatic contracted to be interface //
func (person Person) getName() string {
	return person.Name
}

// Automatic contracted to be interface //
func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	var adam = Person{
		Name: "Adam Syarif Hidayatullah",
	}
	sayHello(adam)
	fmt.Println(adam.getName())

	var anjing = Animal{
		Name: "Anjing",
	}
	sayHello(anjing)
	fmt.Println(anjing.getName())
}
