package main

import (
	"fmt"
	"strings"
)

// Buatlah program yang menghasilkan output seperti dibawah ini.
// "abcd" -> "A-Bb-Ccc-Dddd"
// "RqaEzty" -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// "cwAt" -> "C-Ww-Aaa-Tttt"

func barisHuruf(huruf string) string{
	arr := strings.Split(huruf, "")
	var arrHasil []string

	for i:=0; i<len(huruf); i++{ //c
		for j:=0; j<=i; j++{ //i=1 j=1<=1
			arrHasil = append(arrHasil, arr[i])
		}
		if i==len(huruf)-1 {
			break
		}
		arrHasil = append(arrHasil, "-")
	}

	tostr := strings.Join(arrHasil, "")
	// titleCase := strings.Title(strings.ToLower(tostr))
	titleCase := strings.ToLower(tostr)
	words := strings.Split(titleCase, "-")
	
	for i, word := range words {
		if len(word) > 0 {
			lastChar := string(word[len(word)-1])
			capitalizedLastChar := strings.ToUpper(lastChar)
			words[i] = word[:len(word)-1] + capitalizedLastChar
		}
	}

	result := strings.Join(words, "-")

	// return titleCase
	return result //R-qQ-aaA-eeeE-zzzzZ-tttttT-yyyyyyY
}

func main(){
	s := "RqaEzty"
	hasil := barisHuruf(s)

	fmt.Println(hasil)
}