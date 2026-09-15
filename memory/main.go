package main

import "fmt"

// ---------------------------------------------------------
// DATA SEGMENT: Global mutable variables go here.
// Slower to access compared to stack variables.
// ---------------------------------------------------------
var P = 100

// ---------------------------------------------------------
// CODE SEGMENT: Constants and Function Definitions go here.
// Read-only. They are placed in memory directly from the binary.
// ---------------------------------------------------------
const A = 10

// init() runs first. It gets a Stack Frame, executes, and is destroyed.
func init() {
	fmt.Println("init() executed first: Hello")
}

func main() {
	// A Stack Frame is created for main()

	call() // Triggers the creation of a new Stack Frame for call()

	// Looks for 'A'.
	// 1. Not in main() stack frame.
	// 2. Not in Data segment.
	// 3. Found in Code Segment (Constant).
	fmt.Println("Value of Constant A:", A)
}

func call() {
	// A Stack Frame is created for call()

	// FUNCTION EXPRESSION:
	// The function logic is in the Code Segment.
	// The local variable 'add' in this stack frame just stores a REFERENCE to it.
	add := func(x, y int) {

		// A Stack Frame is created for add()
		z := x + y // 'z' is stored locally in the add() Stack Frame
		fmt.Println("Sum:", z)

		// When add() finishes, its Stack Frame is destroyed.
	}

	// Executes add(). Passes 5 and 6 directly into the local Stack Frame.
	add(5, 6)

	// Variable Lookup in action:
	// 'P' is not local -> fetches 100 from Data Segment (Global)
	// 'A' is not local -> fetches 10 from Code Segment (Constant)
	add(P, A)
}
