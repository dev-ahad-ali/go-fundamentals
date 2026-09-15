package main

import (
	"fmt"
	"unsafe"
)

// callerFunction demonstrates how a function frame holds local variables
func callerFunction() {
	var a int64 = 100
	var b int64 = 200

	fmt.Println("==================================================")
	fmt.Println(" 1. CALLER STACK FRAME MEMORY LAYOUT")
	fmt.Println("==================================================")
	fmt.Printf("Caller Variable 'a' Address : %p (Size: %d bytes)\n", &a, unsafe.Sizeof(a))
	fmt.Printf("Caller Variable 'b' Address : %p (Size: %d bytes)\n", &b, unsafe.Sizeof(b))

	// Invoking child function to trigger nested stack frame creation
	childFunction(a, b)
}

// childFunction receives parameters and allocates its own local variables on a new stack frame
func childFunction(param1, param2 int64) {
	var x int64 = 300
	var y int64 = 400

	fmt.Println("\n==================================================")
	fmt.Println(" 2. CHILD STACK FRAME MEMORY LAYOUT (SP/BP SHIFT)")
	fmt.Println("==================================================")
	// Notice how parameters and local variables reside at different memory addresses
	fmt.Printf("Child Parameter 'param1' Address : %p\n", &param1)
	fmt.Printf("Child Parameter 'param2' Address : %p\n", &param2)
	fmt.Printf("Child Local Var 'x' Address      : %p\n", &x)
	fmt.Printf("Child Local Var 'y' Address      : %p\n", &y)

	// Demonstrating stack memory layout distance
	ptrX := uintptr(unsafe.Pointer(&x))
	ptrY := uintptr(unsafe.Pointer(&y))

	var diff uintptr
	if ptrX > ptrY {
		diff = ptrX - ptrY
	} else {
		diff = ptrY - ptrX
	}

	fmt.Printf("\nOffset between adjacent stack variables (x and y): %d bytes\n", diff)
}

func main() {
	fmt.Println("Starting Stack Memory Trace...")
	callerFunction()
}
