package main

import "fmt"

/**
Interface Kosong is Super Class of all Classes
Interface Kosong is data type
Interface Kosong can be all data type -> but all data type isn't Interface Kosong
Interface Kosong is contract of all contract

ex :
1. fmt.Println(param: a ...interface{})
2. panic(v interface{})
3. recover() interface{} -> return value is interface kosong folowwed panic return value
4. etc
*/

func cobaInterfaceKosong(i int) interface{} {
	if i == 1 {
		return i
	} else if i == 2 {
		return true
	} else {
		return "Benar"
	}
}

func main() {
	// Because if u was knew output will be integer
	// it doesn't mean to be integer. it actually interface kosong
	// but interface kosong can be integer or any data type
	// because interface kosong is super class of any classes
	var interfaceKosong interface{} = cobaInterfaceKosong(1)
	fmt.Println(interfaceKosong)
}
