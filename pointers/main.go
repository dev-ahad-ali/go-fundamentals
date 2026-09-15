package main

import "fmt"

// Employee defines a custom data structure
type Employee struct {
	Name   string
	Salary int
}

// modifyStructByValue receives a completely independent copy of the struct.
// Changes made inside this function will not affect the original instance.
func modifyStructByValue(emp Employee) {
	emp.Salary = 90000
	emp.Name = "Modified Copy"
}

// modifyStructByPointer receives the exact memory address of the struct.
// Changes made here update the original data structure directly.
func modifyStructByPointer(empPtr *Employee) {
	// Go automatically dereferences the pointer under the hood.
	// We do not need to write (*empPtr).Salary
	empPtr.Salary = 90000
	empPtr.Name = "Habib (Promoted)"
}

func main() {
	// 1. Basic Primitive Pointer Operations
	x := 20
	var p *int = &x

	fmt.Println("Initial value of x:", x)
	fmt.Println("Memory address of x:", p)

	*p = 50
	fmt.Println("Value of x after pointer mutation (*p = 50):", x)

	fmt.Println("--------------------------------------------------")

	// 2. Struct Operations: Pass-by-Value vs. Pass-by-Pointer
	// Initialize a new Employee struct instance
	emp1 := Employee{
		Name:   "Habib",
		Salary: 50000,
	}

	fmt.Println("Original Struct State:")
	fmt.Printf("Name: %s, Salary: %d\n", emp1.Name, emp1.Salary)
	fmt.Println("Memory address of emp1:", &emp1)

	fmt.Println("--------------------------------------------------")

	// Attempting to modify the struct by passing it by value
	modifyStructByValue(emp1)
	fmt.Println("Struct State after modifyStructByValue(emp1):")
	fmt.Printf("Name: %s, Salary: %d\n", emp1.Name, emp1.Salary) // Unchanged

	fmt.Println("--------------------------------------------------")

	// Successfully modifying the struct by passing its memory address
	modifyStructByPointer(&emp1)
	fmt.Println("Struct State after modifyStructByPointer(&emp1):")
	fmt.Printf("Name: %s, Salary: %d\n", emp1.Name, emp1.Salary) // Successfully changed
}
