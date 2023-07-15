package main

import (
	"fmt"
)

func factorial(number int) int{
	if number == 1{
		return 1
	}

	factorialOfNumber := number * factorial(number-1)

	return factorialOfNumber
}

func main(){
	var number int

	for{
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)
		if number > 0 {break}
		fmt.Print("Please enter a number!\n")
	}
	fmt.Print(factorial(number))
}