package main

func conditions() {
	// If-Else Statement
	num := 10
	if num%2 == 0 {
		println(num, "is even")
	} else {
		println(num, "is odd")
	}

	// If-Else If-Else Statement
	score := 85
	if score >= 90 {
		println("Grade: A")
	} else if score >= 80 {
		println("Grade: B")
	} else if score >= 70 {
		println("Grade: C")
	} else if score >= 60 {
		println("Grade: D")
	} else {
		println("Grade: F")
	}

	// Switch Statement
	day := 3
	switch day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("Invalid day")
	}

	// Switch with multiple cases
	fruit := "apple"
	switch fruit {
	case "apple", "banana", "cherry":
		println(fruit, "is a common fruit")
	case "kiwi", "mango":
		println(fruit, "is an exotic fruit")
	default:
		println("Unknown fruit")
	}

	// Type Switch
	var value interface{} = 42
	switch v := value.(type) {
	case int:
		println("Integer:", v)
	case string:
		println("String:", v)
	default:
		println("Unknown type")
	}
}
