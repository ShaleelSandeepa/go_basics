package main

import "fmt"

func main() {

	var num int = 42
	fmt.Println("Value of num:", num)

	var ptr *int = &num // pointer to an integer variable
	fmt.Println("Value of ptr (address of num):", ptr)
	fmt.Println("Value at the address stored in ptr:", *ptr)

	*ptr = 100 // changing the value at the address
	fmt.Println("New value of num after changing via ptr:", num)

	// Demonstrating nil pointer
	var nilPtr *int
	fmt.Println("Value of nilPtr (address of nil):", nilPtr)

	// address of the ptr
	fmt.Println("Address of ptr variable:", &ptr)

	// **ptr2 gives the value at the address stored in ptr
	var ptr2 **int = &ptr
	fmt.Println("Value of ptr2 (address of ptr):", ptr2)
	fmt.Println("Value at the address stored in ptr2 (which is ptr):", *ptr2)
	fmt.Println("Value at the address stored in ptr (which is num):", **ptr2)
}
