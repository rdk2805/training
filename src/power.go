package main

import "fmt"

func main() {
	x := 1
	y := 0
	fmt.Printf("Enter a number:")
	fmt.Scan(&x)
	fmt.Printf("Enter an index:")
	fmt.Scan(&y)
	answer := 1
	for i := 0; i < y; i++ {
		answer *= x
	}

	fmt.Println(x, "to the power of", y, "=", answer)
}
