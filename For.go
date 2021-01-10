package main

import "fmt"

func main() {
	//Pengulangan hanya for loop di golang
	var counter = 1 // counter := 1

	for counter <= 10 {
		fmt.Println("Adam", counter)
		counter++
	}

	// init statement (awal -> satu kali); condition statement; post statement (dijalankan diakhir berulang)
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Adam", counter)
	}

	//for range -> array, slice & map
	var dataArr = [...]string{"Adam", "Rekol", "Nana"}                                 // Array [...] atau [5]
	var dataSlc = []string{"Adam", "Rekol", "Nana"}                                    //Slice
	var dataMap = map[string]string{"nama": "Adam", "nama2": "Rekol", "nama3": "Nana"} // Map
	//Cara 1 Declared
	for key, value := range dataMap {
		fmt.Println("key", key, "=", value)
	}
	//Cara 2 Declared but not Used
	for _, value := range dataMap {
		fmt.Println(value)
	}

	fmt.Println(len(dataArr))
	fmt.Println(cap(dataSlc))
}
