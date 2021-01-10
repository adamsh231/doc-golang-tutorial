package main

import (
	"errors"
	"fmt"
)

/**
Error is function contracted by interface
	->  type error interface{
			Error() string
		}
Error has helper package so doesnt have to manually create function
*/

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Nilai pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var contohError error
	contohError = errors.New("Ups Error!")
	fmt.Println(contohError.Error())

	var hasilBagi, hasilError = pembagian(12, 0)
	fmt.Println(hasilBagi)
	fmt.Println(hasilError)
}
