package main

import (
	"fmt"
)

// Buatlah program yang menerima dua input berupa int dan outputnya berupa list bertipe int. 
// asumsikan input pertama adalah s dan input kedua adalah n. 
// Tugas kalian adalah mengalikan s dengan i sebanyak n kali.
// countBy(1,10)  should return  {1,2,3,4,5,6,7,8,9,10}
// countBy(2,5)  should return {2,4,6,8,10}

func countBy(s, n int) []int{
	var list []int
	for i := 1; i <= n; i++ {
		list = append(list, s*i)
	}

	return list
}

func main(){
	output := countBy(1,10)
	output2 := countBy(2,5)

	fmt.Println(output)
	fmt.Println(output2)
}