package main

import "fmt"

func main()  {
	var operator, angka1, angka2 int

	loop:
	fmt.Println("Pilih Operator:")
	fmt.Println("1.+ \n2.- \n3.x \n4./ \n5.%\n")
	fmt.Scanln(&operator)
	if operator<1 || operator>5 {
		goto loop
	}
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)

	tambah, kurang, kali, bagi, modulus := (angka1 + angka2), (angka1 - angka2), (angka1 * angka2), (angka1 / angka2), (angka1 % angka2)
	switch operator {
	case 1: 
		fmt.Print(angka1," + ",angka2, " = ", tambah)
	case 2: 
		fmt.Print(angka1," - ",angka2, " = ", kurang)
	case 3: 
		fmt.Print(angka1," x ",angka2, " = ", kali)
	case 4: 
		fmt.Print(angka1," / ",angka2, " = ", bagi)
	case 5: 
		fmt.Print(angka1," % ",angka2, " = ", modulus)
	}
}