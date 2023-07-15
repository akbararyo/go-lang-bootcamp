package main

import (
	"fmt"
	"strconv"
)

func factorial(number int) int{
	if number == 1{
		return 1
	} else if number <= 0{
		return 0
	}

	factorialOfNumber := number * factorial(number-1)

	return factorialOfNumber
}

func main(){
	var input string
	var number int
	var err error

	for{
		fmt.Print("Enter a number: ")
		_, err = fmt.Scan(&input)
		if err != nil{
			fmt.Print("Please enter a number!\n")
			continue
		}
		number, err = strconv.Atoi(input)
		if err != nil {
			fmt.Print("Please enter a number!\n")
			continue
		}
		break
	}
	fmt.Print(factorial(number))
}