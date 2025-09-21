package main

import "fmt"

// Constants
const piValue float64 = 3.14159

func main() {
	var name string = "Shaleel Sandeepa"
	fmt.Println(name)

	var firstNumber int = 10
	var secondNumber int = 20

	var sum int = firstNumber + secondNumber
	fmt.Println("Sum:", sum)

	var isActive bool = true
	fmt.Println("Is Active:", isActive)

	var score float32 = 95.5
	fmt.Println("Score:", score)

	var pi float64 = 3.14159
	fmt.Println("Value of Pi:", pi)

	var char rune = 'A'
	fmt.Println("Character:", char)

	fmt.Println("Dividing two numbers:", firstNumber/secondNumber)
	fmt.Println("Dividing two numbers:", float64(firstNumber)/float64(secondNumber))

	var division float64 = float64(secondNumber) / float64(firstNumber)
	fmt.Println("Division:", division)
	fmt.Printf("%T\n", division)
	fmt.Println("Type of division variable:", fmt.Sprintf("%T", division))

	// Multiple variable declarations
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// Short variable declaration
	newUserName := "Shaleel"
	fmt.Println(newUserName)
	fmt.Printf("%T\n", newUserName)

	marks := 82
	fmt.Println(marks)
	fmt.Printf("%T\n", marks)

	isPassed := true
	fmt.Println(isPassed)
	fmt.Printf("%T\n", isPassed)

	// Multiple short variable declaration
	x, y, z := 4, 5.6, "Hello"
	fmt.Println(x, y, z)
	fmt.Printf("%T %T %T\n", x, y, z)

	// Blank identifier
	_, age := "Sandeepa", 21
	fmt.Println(age)

	// piValue = 3.14159 // This will cause a compile-time error
	fmt.Println("Constant Pi Value:", piValue)

	// Typed constants
	const typedInt int = 100
	const typedFloat float64 = 99.99
	const typedString string = "Constant String"
	const typedBool bool = false

	// Calculating area of a circle
	var radius float64 = 5
	area := piValue * radius * radius
	fmt.Println("Area of Circle:", area)

	radius = 10
	area = piValue * radius * radius
	fmt.Println("Area of Circle with new radius:", area)

	radius = 7.5
	area = piValue * radius * radius
	fmt.Println("Area of Circle with another radius:", area)
	fmt.Printf("Area of Circle with another radius: %.2f", area)
}
