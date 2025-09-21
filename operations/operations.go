package main

import "fmt"

func main() {
	var intNum int = 42
	var floatNum float64 = 3.14

	sum := intNum + int(floatNum) // Type conversion
	fmt.Println("Sum of intNum and floatNum (after conversion):", sum)

	sum2 := float64(intNum) + floatNum // Type conversion
	fmt.Println("Sum of intNum and floatNum (after conversion):", sum2)

	a := 10
	fmt.Println("Value of a:", a)

	// a := 12.5 // This will cause a compile-time error
	// fmt.Println("Value of a:", a)

	value := intNum * int(floatNum)
	fmt.Println("Value of intNum * floatNum (after conversion):", value)
}
