package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Buatlah program yang menerima input berupa string dan gantilah semua digit di bawah 5 dengan "0", 
// sedangkan digit 5 keatas di ganti dengan "1". Outpu dari program tersebut berupa string

func digitConverter(input string) string{
	var convertedDigits []string
	inputArray := strings.Split(input, "") //jika tidak dibuat array maka tiap digit akan dihitung value stringnya (2=50) bukan valuenya asli

	for i, output := range inputArray{
		a, err := strconv.Atoi(inputArray[i]) //pengecekan di dalam for karena untuk mengecek tiap index
		if err != nil {
			fmt.Println("Masukkan ANGKA!!")
			return "0"
		}else {
			if a<5{
				output = "0"
				convertedDigits = append(convertedDigits, output)
			}else{
				output = "1"
				convertedDigits = append(convertedDigits, output)
			}
		}
		fmt.Printf("%s -> %s\n", inputArray[i], output)
	}
	return strings.Join(convertedDigits, "")
}

func main(){
	var input, output string

	loop:
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&input)
	output = digitConverter(input)
	if output=="0"{goto loop}
	
	fmt.Printf("Hasil konversi dari angka %s yaitu %s", input, output)
}