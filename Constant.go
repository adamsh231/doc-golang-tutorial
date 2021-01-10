package main

import "fmt"

func main() {
	const final = "adam"
	// Tidak wajib dipakai
	const (
		final1 = "adam"
		final2 = "adam"
	)
	fmt.Println(final)
}
