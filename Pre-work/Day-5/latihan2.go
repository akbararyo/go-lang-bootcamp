package main

import (
	"fmt"
	"strings"
)

// Buatlah program yang menkonversi DNA ke RNA. 
// input program berupa string (DNA) dan outputnya berupa string (RNA).
// "GCAT"  =>  "GCAU"

func dnaToRna(dna string) string{
	var rnaArray []string
	dnaUpper := strings.ToUpper(dna)
	dnaArray := strings.Split(dnaUpper, "")

	for _, rna := range dnaArray{
		if rna == "C" || rna == "U" || rna == "A" || rna == "G" {
			if rna == "T"{
				rna = "U"
				rnaArray = append(rnaArray, rna)
			}else {
				rnaArray = append(rnaArray, rna)
			}
		}else {
			// fmt.Println("Masukkan kombinasi DNA!")
			// return ""
		}
		
	}
	return strings.Join(rnaArray,"")
}

func main(){
	input := "GCAT"
	input2 := "ATCGATCG"
	input3 := "ABCDEFG"

	output := dnaToRna(input)
	output2 := dnaToRna(input2)
	output3 := dnaToRna(input3)

	fmt.Printf("DNA=>RNA: %s => %s\n", input, output)
	fmt.Printf("DNA=>RNA: %s => %s\n", input2, output2)
	fmt.Printf("DNA=>RNA: %s => %s\n", input3, output3)
}