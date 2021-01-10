package main

import "fmt"

/**
Type Assertion usually used to handle interface kosong
*/

func random() interface{} {
	return "Ok!"
}

func main() {
	var randomData = random()
	var typeString = randomData.(string) //! Please be sure data type is right because can be panic

	fmt.Println(typeString)

	// use switch case for handle panic type assertion -> unique used switch for type assertion
	switch switchRandom := randomData.(type) {
	case string:
		fmt.Println("String", switchRandom)
	case bool:
		fmt.Println("Boolean", switchRandom)
	case int:
		fmt.Println("Integer", switchRandom)
	case float32:
		fmt.Println("Float", switchRandom)
	default:
		fmt.Println("Unknown")
	}
}
