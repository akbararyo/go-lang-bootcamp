package main

import "fmt"

func contoh1(){
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}

func contoh2(){
	var point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 3:
		fmt.Println("you need to learn more")
	case point < 4:
		fmt.Println("you need to learn more!!")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}

func main(){
	fmt.Println("Contoh 1")
	contoh1()
	fmt.Println("\nContoh 2")
	contoh2()
}