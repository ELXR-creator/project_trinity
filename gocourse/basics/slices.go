package main

// slices are built on top of arrays and provide a more flexible and powerful way to work with sequences of elementsq.
// slices are reference types in go this means when you assign a slice to a new variable or pass it to a function, no copy of the underlying array is made, both slices point to the same array.
func slices() {
	// Creating a slice using a slice literal
	numbers := []int{1, 2, 3, 4, 5}

	//create  an empty slice
	var emptySlice []int

	_ = emptySlice
	// Appending elements to a slice
	numbers = append(numbers, 6, 7, 8)

	// Slicing a slice
	subSlice := numbers[2:5] // elements from index 2 to 4
	_ = subSlice

	// Length and capacity of a slice
	length := len(numbers)   // length of the slice
	capacity := cap(numbers) // capacity of the slice
	_, _ = length, capacity

	// Creating a slice using the make function
	moreNumbers := make([]int, 3) // creates a slice of length 3
	moreNumbers[0] = 10
	moreNumbers[1] = 20
	moreNumbers[2] = 30

	//multidimensional slices
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	_ = matrix
}
