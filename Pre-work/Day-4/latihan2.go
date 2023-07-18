package main

import (
	"fmt"
	"strings"
	"strconv"
)

// Buatlah program yang menentukan abad dari tahun yang di berikan.
// 1705 --> 18
// 1900 --> 19
// 1601 --> 17
// 2000 --> 20

func main(){
	var tahun int

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)

	tahun100 := tahun+100

	sthn, sthn100 := strconv.Itoa(tahun), strconv.Itoa(tahun100)

	arr := strings.Split(sthn,"")
	arrb := strings.Split(sthn100,"")

	abd := strings.Join(arr[:2],"")
	abd100 := strings.Join(arrb[:2],"")

	if tahun%100 == 0 {
		fmt.Print("sekarang abad ke-", abd)
	}else if tahun%100 != 0 {
		fmt.Print("sekarang abad ke-", abd100)
	}
}