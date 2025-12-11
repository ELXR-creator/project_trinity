package main

func loops() {
	// For loop example
	for i := 0; i < 5; i++ {
		// Print the current value of i
		println("For loop iteration:", i)
	}

	// While loop example using for
	j := 0
	for j < 5 {
		// Print the current value of j
		println("While loop iteration:", j)
		j++
	}

	// Infinite loop example (commented out to prevent actual infinite loop)
	/*
		for {
			// This will run indefinitely
			println("Infinite loop iteration")
		}
	*/

	// Looping over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		println("Index:", index, "Value:", value)
	}
}
