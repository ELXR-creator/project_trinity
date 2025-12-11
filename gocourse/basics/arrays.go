package main

// arrays are vakue types in go this means when you assign an array to a new variable or pass it to a function, a copy of the array is made.
// This behavior is different from slices, which are reference types.
func arrays() {
	// Declare and initialize an array
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	println("Array:", arr)

	// Access array elements
	println("First element:", arr[0])
	println("Second element:", arr[1])
	println("Third element:", arr[2])

	// Modify array elements
	arr[2] = 10
	println("Modified third element:", arr[2])

	// Get the length of the array
	length := len(arr)
	println("Length of array:", length)

	// Iterate over the array using a for loop
	println("Array elements:")
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	// Iterate over the array using range
	println("Array elements using range:")
	for index, value := range arr {
		println("Index:", index, "Value:", value)
	}

	// Multi-dimensional array
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	println("Multi-dimensional array (matrix):")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			print(matrix[i][j], " ")
		}
		println()
	}

	// Array comparison
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{4, 5, 6}

	println("arr1 == arr2:", arr1 == arr2) // true
	println("arr1 == arr3:", arr1 == arr3) // false

	//blank identifier to ignore index or value
	println("Using blank identifier to ignore index:")
	for _, value := range arr {
		println("Value:", value)
	}
}
