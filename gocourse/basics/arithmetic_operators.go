package main

import "fmt"

func arithmeticOperators() {
	//Variable declaration
	var a, b int = 15, 4

	//Basic Arithmetic Operators
	//addition +
	var sum int = a + b
	fmt.Println("Addition:", sum)
	//subtraction -
	var difference int = a - b
	fmt.Println("Subtraction:", difference)
	//multiplication *
	var product int = a * b
	fmt.Println("Multiplication:", product)
	//division /
	var quotient int = a / b
	fmt.Println("Division:", quotient)
	//Remainder modulus %
	var remainder int = a % b
	fmt.Println("Remainder:", remainder)
	//Increment ++
	a++
	fmt.Println("Incremented a:", a)
	//Decrement --
	b--
	fmt.Println("Decremented b:", b)

	//Compound Assignment Operators
	//Addition assignment +=
	a += 5
	fmt.Println("After += 5, a:", a)
	//Subtraction assignment -=
	b -= 2
	fmt.Println("After -= 2, b:", b)
	//Multiplication assignment *=
	a *= 2
	fmt.Println("After *= 2, a:", a)
	//Division assignment /=
	b /= 1
	fmt.Println("After /= 1, b:", b)
	//Modulus assignment %=
	a %= 4
	fmt.Println("After %= 4, a:", a)

	//Additional Notes:
	const p float64 = 22 / 7
	fmt.Println("Constant p (22/7):", p)

	const z float64 = 22.0 / 7.0
	fmt.Println("Constant z (22.0/7.0):", z)

	//Integer vs. Floating-Point Division
	var intDiv int = 7 / 2
	fmt.Println("Integer Division 7 / 2:", intDiv)
	var floatDiv float64 = 7.0 / 2.0
	fmt.Println("Floating-Point Division 7.0 / 2.0:", floatDiv)

	//Type Conversion
	var intVal int = 10
	var floatVal float64 = float64(intVal) / 3.0
	fmt.Println("Type Conversion int to float64:", floatVal)

	//Operator Precedence
	//Parentheses ()
	//Multiplication and Division *, /
	//Addition and Subtraction +, -

	//Overflow
	//Underflow

	//Why be Mindful of overflow and underflow?
	//Because they can cause unexpected behavior and errors in your programs.
	// Program Stability
	//Security Vulnerabilities
	//Data Integrity
	//Type Safety

	//Mitigation Strategies
	//Range Checking
	//Type Conversion
	//Error Handling

}
func main() {
	arithmeticOperators()
}
