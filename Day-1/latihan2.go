package main

import "fmt"

func main(){
	var nama string

	fmt.Print("Nama kamu siapa? ")
	fmt.Scan(&nama)
	fmt.Printf("Selamat siang, %s", nama)
}