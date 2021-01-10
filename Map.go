package main

import "fmt"

func main() {
	// map[key-data-type]value-data-type
	var person = map[string]string{
		"nama":  "Adam Syarif Hidayatullah",
		"Idola": "Nabi Muhammad SAW",
	}
	person["pekerjaan"] = "programmer"
	fmt.Println(person)
	fmt.Println(person["nama"])

	// make function
	var person2 = make(map[string]string)
	person2["pekerjaan"] = "programmer"
	fmt.Println(person2)
	delete(person2, "pekerjaan")
	fmt.Println(person2)
}
