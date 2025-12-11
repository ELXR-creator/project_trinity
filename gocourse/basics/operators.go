package main

func operators() {
	// Arithmetic Operators
	a := 10
	b := 3
	println("Addition:", a+b)       // 13
	println("Subtraction:", a-b)    // 7
	println("Multiplication:", a*b) // 30
	println("Division:", a/b)       // 3
	println("Modulus:", a%b)        // 1
	println("Increment:", a+1)      // 11
	println("Decrement:", b-1)      // 2

	// Relational Operators
	println("Equal:", a == b)                 // false
	println("Not Equal:", a != b)             // true
	println("Greater Than:", a > b)           // true
	println("Less Than:", a < b)              // false
	println("Greater Than or Equal:", a >= b) // true
	println("Less Than or Equal:", a <= b)    // false

	// Logical Operators
	x := true
	y := false
	println("Logical AND:", x && y) // false
	println("Logical OR:", x || y)  // true
	println("Logical NOT:", !x)     // false

	// Bitwise Operators
	c := 5                            // 0101 in binary
	d := 3                            // 0011 in binary
	println("Bitwise AND:", c&d)      // 1 (0001 in binary)
	println("Bitwise OR:", c|d)       // 7 (0111 in binary)
	println("Bitwise XOR:", c^d)      // 6 (0110 in binary)
	println("Bitwise AND NOT:", c&^d) // 4 (0100 in binary)
	println("Left Shift:", c<<1)      // 10 (1010 in binary)
	println("Right Shift:", c>>1)     // 2 (0010 in binary)

	//Comparison Operators
	println("c == d:", c == d)
	println("c != d:", c != d)
	println("c > d:", c > d)
	println("c < d:", c < d)
	println("c >= d:", c >= d)
	println("c <= d:", c <= d)

	// Assignment Operators
	e := 10
	e += 5                   // e = e + 5
	println("After += :", e) // 15
	e -= 3                   // e = e - 3
	println("After -= :", e) // 12
	e *= 2                   // e = e * 2
	println("After *= :", e) // 24
	e /= 4                   // e = e / 4
	println("After /= :", e) // 6
	e %= 4                   // e = e % 4
	println("After %= :", e) // 2

	// Miscellaneous Operators
	f := 20
	println("Address of f:", &f) // Address of f
	println("Value of f:", *&f)  // 20

	var g interface{} = "Hello, Go!"
	str, ok := g.(string)
	if ok {
		println("Type Assertion successful:", str) // Hello, Go!
	} else {
		println("Type Assertion failed")
	}

	var h interface{} = 42
	num, ok := h.(int)
	if ok {
		println("Type Assertion successful:", num) // 42
	} else {
		println("Type Assertion failed")
	}

}
