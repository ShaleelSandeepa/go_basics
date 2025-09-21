package main

import "fmt"

func main() {
	a := 10
	b := 3

	sum := a + b
	diff := a - b
	prod := a * b
	quot := a / b // division operator | integer division | because both are integers
	mod := a % b  // modulus operator | remainder

	println("Sum:", sum)
	println("Difference:", diff)
	println("Product:", prod)
	println("Quotient:", quot)
	println("Modulus:", mod)

	// For float division, we need to convert at least one operand to float
	quotFloat := float64(a) / float64(b) // Uncomment this line for float division
	println("Float Quotient:", quotFloat)

	// Increment and Decrement
	c := 5
	fmt.Println("Before incrementing, c =", c)
	c++
	fmt.Println("After incrementing, c =", c)
	c--
	fmt.Println("After decrementing, c =", c)

	// Using compound assignment operators
	d := 5
	d += 2
	fmt.Println("After adding 2, d =", d)
	d *= 3
	fmt.Println("After multiplying by 3, d =", d)
	d -= 4
	fmt.Println("After subtracting 4, d =", d)
	d /= 2
	fmt.Println("After dividing by 2, d =", d)
	d %= 3
	fmt.Println("After modulus by 3, d =", d)

	// Demonstrating operator precedence
	result := a + b*2 - (a / b) + (a % b)
	fmt.Println("Result of a + b*2 - (a/b) + (a % b) =", result)

	// Bitwise operations
	x := 6 // 110 in binary
	y := 3 // 011 in binary

	fmt.Println("Bitwise AND:", x&y)  // 010 in binary
	fmt.Println("Bitwise OR:", x|y)   // 111 in binary
	fmt.Println("Bitwise XOR:", x^y)  // 101 in binary
	fmt.Println("Bitwise NOT:", ^x)   // 001 in binary (inverted bits)
	fmt.Println("Left Shift:", x<<1)  // 1100 in binary
	fmt.Println("Right Shift:", x>>1) // 011 in binary

	// Logical operations
	p := true
	q := false
	fmt.Println("Logical AND:", p && q)
	fmt.Println("Logical OR:", p || q)
	fmt.Println("Logical NOT:", !p)
}
