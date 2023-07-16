package main

import (
	"fmt"
	"strings"
)

// Buatlah program yang membalikan semua kata dalam sebuah string. 
// Setiap kata di pisahkan dengan tepat satu spasi dan tidak ada spasi di awal ataupun di akhir string.

// "The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The"

func reverseString(original []string){
	for i:=len(original)-1; i>=0; i--{
        fmt.Print(original[i], " ")
    }
}

func main(){
	original := "The greatest victory is that which requires no battle"
	arr := strings.Split(original," ")
	
	reverseString(arr)
}