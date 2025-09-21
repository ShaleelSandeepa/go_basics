package main

import "fmt"

func main() {

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
	fmt.Println("First Element:", arr[0])
	fmt.Println("Length of Array:", len(arr))
	fmt.Printf("Type of arr: %T\n", arr)

	// Iterating over the array
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element at index %d: %d\n", i, arr[i])
	}

	var colors [3]string
	fmt.Println("Array of strings:", colors)
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println("Colors Array:", colors)

	// Multidimensional array
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D Array (Matrix):", matrix)
	fmt.Println("Element at (1,2):", matrix[1][2])

	// Array with inferred size
	numbers := [...]int{10, 20, 30, 40}
	fmt.Println("Inferred Size Array:", numbers)
	fmt.Println("Length of numbers array:", len(numbers))
	numbers[2] = 100
	fmt.Println("Updated numbers array:", numbers)

	// Multi-dimensional array with inferred size
	grid := [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println("Inferred Size 2D Array:", grid)
	fmt.Println("Element at (2,1):", grid[2][1])

	// Multi-dimensional array with different data types
	var mixedArray [2][2]interface{} = [2][2]interface{}{ // using interface{} to hold different types
		{1, "Hello"},
		{3.14, true},
	}
	fmt.Println("Mixed Data Type 2D Array:", mixedArray)
	fmt.Println("Element at (0,1):", mixedArray[0][1])
	mixedArray[1][0] = "Changed"
	fmt.Println("Updated Mixed Data Type 2D Array:", mixedArray)

	var mixedArray2 [2][2]any = [2][2]any{ // using any to hold different types
		{1, "Hello"},
		{3.14, true},
	}
	fmt.Println("Mixed Data Type 2D Array:", mixedArray2)
	fmt.Println("Element at (0,1):", mixedArray2[0][1])
	mixedArray2[1][0] = "Changed"
	fmt.Println("Updated Mixed Data Type 2D Array:", mixedArray2)
}
