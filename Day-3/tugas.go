package main

import "fmt"

// Buatlah sebuah program yang berfunsi untuk meringkas sebuah array bertipe integer (n). 
// Setiap elemen pada array tersebut di bandingkan dengan elemen selanjutnya. 
// Jika hasil perkalian antara elemen tersebut kurang dari atau sama dengan (k) maka gabungkan elemen tersebut. 
// Hasil akhirnya kembalikan panjang dari array yang telah di ringkas tersebut.

// n = [1,2,4,7,1]
// k = 9

// Iterasi-1
// [2 4 7 1]

// Iterasi-2
// [8 7 1]

// Iterasi-3
// [8 7]

// Hasil = 2

func banding(n []int, k int) (hasil int, arr []int){
	for i:=0; i<=len(n)-2; i++{
		switch{
		case (n[i]*n[i+1])<=k:
			n[i] = n[i]*n[i+1]
			n = append(n[:i+1], n[i+2:]...)
			i--
			fmt.Printf("%d\n", n)
		}
	}
	hasil = len(n)
	arr = n
	return
}

func main(){
	n := []int{1,2,4,7,1}
	k := 9

	hasil, arr := banding(n, k)
	fmt.Printf("Hasil: %d\nPanjang:%d\n", arr, hasil)
}