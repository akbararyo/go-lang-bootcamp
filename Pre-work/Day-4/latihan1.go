package main

import (
	"fmt"
	"strings"
)

// Buatlah program yang menghilangkan karakter pertama dan terakhir pada sebuah string.
// "hello world" => "ello worl"

func hapusAwalAkhir(s string){
	
}

func main(){
	s := "hello world"
	//arr := [...]string{s}
	arr := strings.Split(s, "")

	var sli = arr[1:len(arr)-1]

	tostr := strings.Join(sli,"")

	fmt.Println(tostr)
}