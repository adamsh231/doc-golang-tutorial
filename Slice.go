package main

import "fmt"

func main() {
	// [...] kalau tidak tau panjang pastinya
	var arrSlc = [...]string{
		"Adam1",
		"Adam2",
		"Adam3",
		"Adam4",
		"Adam5",
		"Adam6",
		"Adam7",
		"Adam8",
		"Adam9",
		"Adam10",
		"Adam11",
		"Adam12",
	}

	// Pointer 3, Length 4 (3 s/d 6), Capacity [len_arr - 3]
	var slice = arrSlc[8:]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//! Array dirubah slice pun juga terubah -> sebaliknya juga
	arrSlc[3] = "dirubah"
	fmt.Println(slice)

	//! Note append -> merubah array reference jika capacity penuh ganti array baru!!
	var slice2 = append(slice, "jembas")
	fmt.Println(slice2)
	fmt.Println(arrSlc)

	// Membuat slice lebih aman karena array panjang 10 terprivate
	var makeSlice = make([]string, 2, 10)
	makeSlice[0] = "mahabare"
	makeSlice[1] = "mahabare2"
	fmt.Println(makeSlice)

	// copy slice disamakan agar tidak terpotong
	var copySlice = make([]string, len(makeSlice), cap(makeSlice))
	copy(copySlice, makeSlice)
	fmt.Println(copySlice)

	// perbedaan deklarasi array vs slice
	var iniArray = [...]int{1, 2, 3, 4, 5}
	var iniSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
