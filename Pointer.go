package main

import "fmt"

/**
Go -> Default is Pass by value
	*PBV	-> Array
			-> Struct
	*PBR	-> Slice //! Coz there is function copy for PBV
			-> Map
Pointer -> Pass by reference in memory
Pointer -> Used to save memory, for example
		-> If pointer not used
		-> If big struct data called in function, Memory will duplicate

Note:
	var x = &y -> pointer reference to y
	*x = struct -> any pointer same reference to x will new declared
	*x.field = ... -> any pointer same reference to x will change field value to ...
	var x = &struct{..,..,..} -> new filled pointer
	var x = new(struct) -> new empty pointer

	* -> can be 'operator' for any pointer reference or class pointer like '*PassByValue'
*/

type PassByValue struct {
	City     string
	Province string
	Country  string
}

func functionPBV(pbv PassByValue) {
	pbv.Country = "Negara Berubah"
}

func functionPBR(pbr *PassByValue) {
	pbr.Country = "Negara Berubah"
}

func (pbv PassByValue) structPBV() {
	pbv.Country = "Negara Berubah"
}
func (pbr *PassByValue) structPBR() {
	pbr.Country = "Negara Berubah"
}

func main() {
	// Pass by value
	var pbv = PassByValue{"Malang", "Jawa Timur", "Indonesia"}
	var pbv_copy = pbv
	pbv_copy.City = "Surabaya"
	fmt.Println(pbv)
	fmt.Println(pbv_copy)

	fmt.Println("---------------")

	// Pass by reference -> pointer -> '&'
	var pbr PassByValue = PassByValue{"Malang", "Jawa Timur", "Indonesia"}
	var pbr_copy *PassByValue = &pbr //! * is pointer
	pbr_copy.City = "Surabaya"
	fmt.Println(pbr)
	fmt.Println(pbr_copy)

	fmt.Println("------")

	pbr_copy = &PassByValue{"Ini", "Dirubah", "Semua"} //! Redeclared doesnt change, changed pointer memory
	fmt.Println(pbr)
	fmt.Println(pbr_copy)

	fmt.Println("--------------")

	// Pass by reference -> pointer -> '*'
	var pointer1 PassByValue = PassByValue{"xxx", "yyy", "zzz"}
	var pointer2 *PassByValue = &pointer1
	var pointer3 *PassByValue = &pointer1

	*pointer3 = PassByValue{"Ini", "Dirubah", "Semua"} //! Change all reference pointer

	fmt.Println(pointer1)
	fmt.Println(pointer2)
	fmt.Println(pointer3)

	fmt.Println("--------------")

	// Pass by reference -> pointer -> 'new'
	var newpointer = new(PassByValue)
	var newdeclarepointer *PassByValue = &PassByValue{"xxx", "yyy", "zzz"}
	var change = newdeclarepointer //! Auto use pointer
	change.City = "Malang"
	fmt.Println(newpointer)
	fmt.Println(newdeclarepointer)
	fmt.Println(change)

	fmt.Println("--------------")

	// Pass by reference -> pointer -> function
	var funcPassBV = PassByValue{"xxx", "yyy", "zzz"}
	var funcPassBR = PassByValue{"xxx", "yyy", "zzz"}
	var declaredpointer = &PassByValue{"xxx", "yyy", "zzz"} //! declared pointer

	functionPBV(funcPassBV)
	functionPBR(&funcPassBR)
	functionPBR(declaredpointer)

	fmt.Println(funcPassBV)
	fmt.Println(funcPassBR)
	fmt.Println(declaredpointer)

	fmt.Println("--------------")

	// Pass by reference -> pointer -> struct function
	var funcPassBV2 = PassByValue{"xxx", "yyy", "zzz"}
	var funcPassBR2 = PassByValue{"xxx", "yyy", "zzz"}

	funcPassBV2.structPBV()
	funcPassBR2.structPBR()

	fmt.Println(funcPassBV2)
	fmt.Println(funcPassBR2)
}
