package main

import "fmt"

// No Param
func sayHello() {
	fmt.Println("Hello World!")
}

// With param
func sayHelloTo(name string) {
	fmt.Println("Hello to my little friend ", name, "!!")
}

// With Return Value
func multiple(value1 int, value2 int) int {
	return value1 * value2
}

// With Return Multiple Value
func getFullName(firstName string, lastName string, ignoredValue []int) (string, string, []int) {
	return firstName, lastName, ignoredValue
}

// With Return Multiple Default Vale
func getDefaultName() (firstName string, lastName string, ignoredValue int) {
	firstName = "Adam"
	ignoredValue = 3
	return
}

// With variadic function -> array behavior but limited
//! Variadic function should in the end of param
func variadic(ignoredValue string, numbers ...int) int {
	var total = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Function as Parameter
func filter(name string) string {
	if name == "nana" {
		return "mantap"
	} else {
		return name
	}
}

type Filter func(string) string  // with alias
type BlackList func(string) bool // with alias

func funcInParam(name string, filter Filter) {
	var filtered = filter(name)
	fmt.Println(filtered)
}

func login(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("Welcome!!")
	} else {
		fmt.Println("Blocked!!")
	}
}

// Recursive Function
func factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

func main() {
	sayHello()

	sayHelloTo("Nana")

	var x = multiple(2, 3)
	fmt.Println(x)

	// ignored with underscores
	var first, lastName, _ = getFullName("Adam", "Syarif", []int{1, 2, 3})
	fmt.Println(first, lastName)

	var first2, lastName2, _ = getDefaultName()
	fmt.Println(first2, lastName2)

	// Variadic
	var total = variadic("Ignored", 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	// slice to param
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var total2 = variadic("Ignored", slice...)
	fmt.Println(total2)

	// save function in variabel
	var variabel = sayHello
	variabel()

	// function in param
	funcInParam("nana", filter)

	var funcfilter = filter
	funcInParam("adam", funcfilter)

	// Anonymous function
	var blacklist = func(name string) bool { //! closure -> beware coz closure can access local data. create new var !!
		if name != "admin" {
			return false
		} else {
			return true
		}
	}

	login("admin", blacklist)

	login("nana", func(name string) bool {
		if name != "admin" {
			return false
		} else {
			return true
		}
	})

	// Recursive function
	fmt.Println(factorial(3))

}
