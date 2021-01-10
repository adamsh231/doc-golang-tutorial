package main

import "fmt"

func main() {
	var name string
	name = "adam"
	fmt.Println(name)

	var name2 = "adam"
	fmt.Println(name2)

	name3 := "adam"
	fmt.Println("nama", name3)

	//variabel wajib dipakai
	var (
		name4 = "adam"
		name5 = "adam"
		name6 = "adam"
	)
	fmt.Println(name4, name5, name6)
}
