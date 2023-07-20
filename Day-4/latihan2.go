package main

import "fmt"

// Buatlah program yang menentukan abad dari tahun yang di berikan.
// 1705 --> 18
// 1900 --> 19
// 1601 --> 17
// 2000 --> 20

func lacakAbad(n int) int{
	var abad int

	if n%100 == 0 {
		abad = n/100
	} else if n%100 != 0 {
		abad = (n+100)/100
	}

	return abad
}

func main(){
	var tahun int

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)

	abad := lacakAbad(tahun)

	fmt.Printf("Tahun %d merupakan abad ke-%d", tahun, abad)
}