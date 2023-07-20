package main

import (
	"fmt"
	"strings"
)

// Buatlah program yang menghilangkan karakter pertama dan terakhir pada sebuah string.
// "hello world" => "ello worl"

func hapusAwalAkhir(s string) string{
	if len(s)<=2{
		return ""
	}else {
		arr := strings.Split(s, "")

		var hapusIndex = arr[1:len(arr)-1]
		tostr := strings.Join(hapusIndex,"")
	
		return tostr
	}
}

func main(){
	s := "hello world"
	s2 := "halo dunia"

	// fmt.Println(s[1:len(s)-1])

	hasil := hapusAwalAkhir(s)
	hasil2 := hapusAwalAkhir(s2)

	fmt.Println(hasil)
	fmt.Println(hasil2)
}