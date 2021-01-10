package main

import "fmt"

func main() {
	type NoKTP string

	var KTPku NoKTP
	KTPku = "3276021310980006"
	fmt.Println(KTPku)
	fmt.Println(NoKTP("1200031324444"))
}
