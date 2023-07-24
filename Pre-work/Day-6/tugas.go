package main

import (
	"fmt"
)

// Tugas
// Buatlah program yang menerima dua input bertipe int (a, b) 
// dan hitunglah jumlah total antara a dan b (dua2 nya termasuk). 
// Jika a dan b bernilai sama maka kembalikan salah satunya.
// CATATAN:$~$a dan b tidaklah berurutan.
// (1, 0) --> 1 (1 + 0 = 1)
// (1, 2) --> 3 (1 + 2 = 3)
// (0, 1) --> 1 (0 + 1 = 1)
// (1, 1) --> 1 (1 since both are same)
// (-1, 0) --> -1 (-1 + 0 = -1)
// (-1, 2) --> 2 (-1 + 0 + 1 + 2 = 2)

func hitung(a, b int) int{
	if a == b {
		return 0
	}else {
		return a + b
	}
}

func main(){
	var a, b int
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&b)
	hasil := hitung(a, b)
	if hasil == 0{
		fmt.Printf("(%d, %d) --> %d since both are the same", a, b, a)
	}else {
		fmt.Printf("(%d, %d) --> %d (%d + %d)", a, b, hasil, a, b)
	}
}