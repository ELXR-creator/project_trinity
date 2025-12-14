package main

func maps() {
	// Creating a map using a map literal
	person := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	//create a map using the make function
	cityPopulation := make(map[string]int)
	cityPopulation["New York"] = 8419600
	cityPopulation["Los Angeles"] = 3980400

	//create an empty map
	var emptyMap map[string]int
	_ = emptyMap

	// Adding or updating key-value pairs
	person["Charlie"] = 35
	person["Alice"] = 31 // update Alice's age

	// Accessing values by key
	ageOfBob := person["Bob"]
	_ = ageOfBob

	// Deleting a key-value pair
	delete(person, "Bob")

	// Checking if a key exists
	age, exists := person["Bob"]
	_ = age
	_ = exists

	// Iterating over a map
	for name, age := range person {
		_ = name
		_ = age
	}

	// Getting the number of key-value pairs in a map
	size := len(person)
	_ = size
}
