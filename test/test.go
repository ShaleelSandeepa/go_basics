package main

import "fmt"

func Increment(n *int) {
	*n++
}

func ConvertToZero(n *int) {
	*n = 0
}

func CheckTheCondition(n *int, isConvertToZero bool) {
	if isConvertToZero {
		ConvertToZero(n)
	} else {
		Increment(n)
	}
}

func main() {
	var number int = 10
	fmt.Println("Initial number:", number)
	CheckTheCondition(&number, false)
	fmt.Println("Output:", number)
}
