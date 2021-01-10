package main

import "fmt"

/**
Nil like Null
Unlike any other language, Go used default value
ex:
Int = 0, Boolean = false, string = "", etc
Nil is Empty data
Nil can be use in function, slice, map, array, pointer, and channel
*/

func newMap(name map[string]string) {
	if name == nil {
		fmt.Println("Data Kosong!")
	} else {
		fmt.Println(name)
	}
}

func main() {
	var person = make(map[string]string) // default is empty
	var person2 map[string]string        // default is nil

	newMap(person)
	newMap(person2)
}
