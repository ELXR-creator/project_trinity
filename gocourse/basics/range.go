package main

func rangeExample() {
	// Using range with a slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		_ = index // index of the element
		_ = value // value of the element
	}

	// Using range with a map
	person := map[string]int{"Alice": 30, "Bob": 25}
	for name, age := range person {
		_ = name // key of the map
		_ = age  // value of the map
	}

	// Using range with a string
	text := "Hello"
	for index, char := range text {
		_ = index // index of the character
		_ = char  // character (as rune)
	}

	// Using range with an array
	array := [5]int{10, 20, 30, 40, 50}
	for index, value := range array {
		_ = index // index of the element
		_ = value // value of the element
	}

	// Using range with a channel (example)
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for value := range ch {
		_ = value // value received from the channel
	}
}
