package main

import "fmt"

func main() {
	x, y := 10, 20

	fmt.Printf("Sebelum:\nx= %d\ny= %d", x, y)

	x, y = y, x

	fmt.Printf("\n\nSesudah:\nx= %d\ny= %d", x, y)
}
