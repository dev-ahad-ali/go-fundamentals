package main

import "fmt"

func main() {
	// Signed Integer with explicit bit allocation (range: -128 to 127)
	var smallNum int8 = -128
	fmt.Printf("int8 value: %d\n", smallNum)

	// Unsigned Integer (only positive numbers and zero, range: 0 to 255)
	var positiveOnly uint8 = 255
	fmt.Printf("uint8 value: %d\n", positiveOnly)

	// Floating-point number restricted to 2 decimal places
	var pi float64 = 3.14159265
	fmt.Printf("Float formatted (2 decimals): %.2f\n", pi)

	// Boolean variable
	var isActive bool = false
	fmt.Printf("Boolean value: %v\n", isActive)

	// Rune type (alias for int32) used for Unicode character representation
	var heartRune rune = '♥'
	fmt.Printf("Rune integer representation: %d\n", heartRune)
	fmt.Printf("Rune character output: %c\n", heartRune)

	// String declaration
	var greeting string = "Hello, Go!"
	fmt.Printf("String output: %s\n", greeting)

	// Runtime Type Inspection using %T
	fmt.Printf("Type of smallNum: %T\n", smallNum)
	fmt.Printf("Type of heartRune: %T\n", heartRune)
	fmt.Printf("Type of isActive: %T\n", isActive)
}
