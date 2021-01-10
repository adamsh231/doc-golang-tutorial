package main

import "fmt"

func main() {
	var arr1 [3]string
	arr1[0] = "adam"
	fmt.Println(arr1)

	// [...] kalau tidak tau panjang pastinya
	var arr2 = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(arr2)
}
