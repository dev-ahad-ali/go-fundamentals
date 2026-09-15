package main

import "fmt"

// Rule 1: Immediate Argument Evaluation vs Deferred Closure
func demoArgumentEvaluation() {
	i := 10

	// Argument i is evaluated immediately at this line.
	// It captures the value 10.
	defer fmt.Println("Deferred with immediate argument evaluation:", i)

	// In a closure without explicit arguments, the variable i is referenced dynamically.
	// It reads the final value of i at execution time.
	defer func() {
		fmt.Println("Deferred closure referencing variable i:", i)
	}()

	i = 20
	fmt.Println("Current value of i in function body:", i)
}

// Rule 2: Execution Order (LIFO)
func demoLIFOOrder() {
	fmt.Println("Registering defers in order: First, Second, Third")

	defer fmt.Println("Deferred 1 (Pushed first, runs last)")
	defer fmt.Println("Deferred 2 (Pushed second, runs second)")
	defer fmt.Println("Deferred 3 (Pushed third, runs first)")
}

// Rule 3 (Part A): Unnamed Return Value
func demoUnnamedReturn() int {
	x := 5
	defer func() {
		// x is modified, but the return value was already copied/frozen as 5
		x = x + 10
	}()
	return x // Returns 5
}

// Rule 3 (Part B): Named Return Value
func demoNamedReturn() (result int) {
	defer func() {
		// Modifies the named return variable 'result' directly before exiting
		result = result + 10
	}()
	result = 5
	return // Returns 15
}

func main() {
	fmt.Println("--- 1. Argument Evaluation ---")
	demoArgumentEvaluation()

	fmt.Println("\n--- 2. LIFO Execution Order ---")
	demoLIFOOrder()

	fmt.Println("\n--- 3. Return Value Behavior ---")
	fmt.Println("Unnamed return result:", demoUnnamedReturn())
	fmt.Println("Named return result:", demoNamedReturn())
}
