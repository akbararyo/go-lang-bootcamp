package main

import (
	"fmt"
	"strings"
)

//Buatlah program yang menerima input berupa string 
// dan mengembalikan panjang dari kata terpendek input tersebut.

func longWordFinder(word string) string{
	wordArray := strings.Split(word, " ")
	var temp string = ""

	for _, this := range wordArray{
		if len(this) > len(temp){
			temp = this
		}
	}
	return temp
}

func main(){
	kalimat := "Buatlah program yang menerima input berupa string dan mengembalikan panjang dari kata terpendek input tersebut."
	longestWord := longWordFinder(kalimat)
	fmt.Print(longestWord)
}