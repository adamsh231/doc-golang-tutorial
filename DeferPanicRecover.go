package main

import "fmt"

//! Panic is stop entire program, so be sure to recover

// Defer Function -> always run while error or not and run in the end
func cobaDefer() {
	var message = recover() // Recover Function
	if message != nil {
		fmt.Println("Panic Error! ,", message) // get panic message
	}

	fmt.Println("Aplikasi Selesai!!")
}

func mainDefer(a int) {
	defer cobaDefer() //! Selalu taruh atas
	fmt.Println("Run Application")
	var result = 10 / a // Error initiate panic function -> defer still running
	fmt.Println(result)
}

// Panic Function -> get Call after error
func appError(isError bool) {
	defer cobaDefer()
	if isError {
		panic("Aplikasi Error!!")
	}
	fmt.Println("Aplikasi Berjalan")
}

// Recover Function -> Handle Panic Error

func main() {
	fmt.Println("----- Defer -----")

	// defer
	mainDefer(0)

	fmt.Println("----- Panic -----")

	// panic
	appError(true)
}
