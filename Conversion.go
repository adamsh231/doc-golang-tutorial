package main

import "fmt"

func main() {
	var integer int32 = 100000
	var integer2 int64 = int64(integer)
	var integer3 int8 = int8(integer2)
	fmt.Println(integer)
	fmt.Println(integer2)
	fmt.Println(integer3)

	var setring = "eko"
	var e = setring[0]
	var x = string(e)
	fmt.Println(x)
}
