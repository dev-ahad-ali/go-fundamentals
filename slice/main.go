package main

import "fmt"

func changeSlice(n []int) []int {
	n[0] = 10
	n = append(n, 11)
	return n
}

// variadic function
func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main() {
	arr := [6]string{"This", "is", "a", "GO", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4] // ["is", "a", "GO"]
	fmt.Println(s)

	s1 := s[1:2] // ["a"] len = 1 cap = 4
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	sl := []int{1, 2, 3} // slice literal
	fmt.Println("Slice:", sl, "len:", len(sl), "capacity:", cap(sl))

	sm := make([]int, 3) // [0, 0, 0], len = 3,  cap = 3
	sm[0] = 5            // [5, 0, 0], len = 3,  cap = 3

	fmt.Println(sm)
	fmt.Println(len(sm))
	fmt.Println(cap(sm))

	sm1 := make([]int, 3, 5) // [0, 0, 0], len = 3, cap = 5
	sm1[0] = 5               // [5, 0, 0], len = 3,  cap = 5
	sm1[2] = 10              // [5, 0, 10], len = 3,  cap = 5

	fmt.Println(sm1)
	fmt.Println(len(sm1))
	fmt.Println(cap(sm1))

	var sn []int             // empty slice or nil slice []
	sn = append(sn, 1, 2, 3) // [1]
	fmt.Println(sn)

	// interview examples
	var x []int      // [], len = 0, cap = 0
	x = append(x, 1) // [1], len = 1, cap = 1
	x = append(x, 2) // [1,2], len = 2 , cap = 2
	x = append(x, 3)

	y := x

	x = append(x, 4)
	y = append(y, 5)

	x[0] = 10

	fmt.Println(x) // [10, 2, 3, 5]
	fmt.Println(y) // [10, 2, 3, 5]

	p := []int{1, 2, 3, 4, 5}
	p = append(p, 6)
	p = append(p, 7)

	l := p[4:]

	q := changeSlice(l)

	fmt.Println(p)      // [1, 2, 3, 4, 10, 6, 7]
	fmt.Println(q)      // [10, 6, 7, 11]
	fmt.Println(p[0:8]) // [1, 2, 3, 4, 10, 6, 7, 11]

	print(4, 53, 5, 3, 2, 5)
}

func init() {
	fmt.Println("This will be invoked first")
}

// detailed example with codes and comments -----

// Variadic function: receives arguments as a slice ([]int) internally
func calculateSum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Modifying an element mutates shared memory
func modifyElement(s []int) {
	if len(s) > 0 {
		s[0] = 999 // Affects caller's underlying array
	}
}

// Appending without a pointer modifies only the local slice header
func appendLocal(s []int) {
	s = append(s, 888) // Caller will NOT see this addition
}

// Appending with a slice pointer modifies the caller's slice header
func appendWithPointer(s *[]int) {
	*s = append(*s, 777) // Updates caller's len/cap and pointer
}

func other() {
	// ---------------------------------------------------------------
	// 1. Array vs. Slice & Header Inspection
	// ---------------------------------------------------------------
	fmt.Println("=== 1. Slice Creation & Header Properties ===")
	// Creating a slice using make(type, len, cap)
	s1 := make([]int, 3, 5)
	fmt.Printf("s1: %v | len: %d | cap: %d\n", s1, len(s1), cap(s1))

	// ---------------------------------------------------------------
	// 2. Slicing & Shared Memory Behavior
	// ---------------------------------------------------------------
	fmt.Println("\n=== 2. Slicing & Shared Memory ===")
	underlyingArray := [5]int{10, 20, 30, 40, 50}
	sliceA := underlyingArray[1:4] // Elements at index 1, 2, 3 -> [20, 30, 40]

	fmt.Println("Original Array: ", underlyingArray)
	fmt.Printf("sliceA: %v | len: %d | cap: %d\n", sliceA, len(sliceA), cap(sliceA))

	// Mutating sliceA alters underlyingArray because they share memory
	sliceA[0] = 222
	fmt.Println("Array after sliceA[0] = 222: ", underlyingArray)

	// ---------------------------------------------------------------
	// 3. Append & Reallocation Trigger
	// ---------------------------------------------------------------
	fmt.Println("\n=== 3. Append & Capacity Reallocation ===")
	s2 := make([]int, 2, 3) // len=2, cap=3
	s2[0] = 1
	s2[1] = 2
	fmt.Printf("Before append (within cap): %v | len: %d | cap: %d\n", s2, len(s2), cap(s2))

	// Appending within capacity (len becomes 3, cap remains 3)
	s2 = append(s2, 3)
	fmt.Printf("After 1st append (len==cap): %v | len: %d | cap: %d\n", s2, len(s2), cap(s2))

	// Appending beyond capacity triggers REALLOCATION (capacity doubles)
	s2 = append(s2, 4)
	fmt.Printf("After 2nd append (reallocated): %v | len: %d | cap: %d\n", s2, len(s2), cap(s2))

	// ---------------------------------------------------------------
	// 4. Function Passing Dynamics
	// ---------------------------------------------------------------
	fmt.Println("\n=== 4. Passing Slices to Functions ===")
	demoSlice := []int{10, 20, 30}
	fmt.Println("Initial demoSlice: ", demoSlice)

	// A. Element modification
	modifyElement(demoSlice)
	fmt.Println("After modifyElement: ", demoSlice) // Element 0 changed to 999

	// B. Local append (fails to update caller)
	appendLocal(demoSlice)
	fmt.Println("After appendLocal: ", demoSlice) // Unchanged length

	// C. Pointer append (successfully updates caller)
	appendWithPointer(&demoSlice)
	fmt.Println("After appendWithPointer: ", demoSlice) // Element 777 added

	// ---------------------------------------------------------------
	// 5. Variadic Functions
	// ---------------------------------------------------------------
	fmt.Println("\n=== 5. Variadic Functions ===")
	// Passing individual arguments
	res1 := calculateSum(10, 20, 30, 40)
	fmt.Println("Sum of individual args:", res1)

	// Unpacking an existing slice into a variadic function using '...'
	numbers := []int{1, 2, 3, 4, 5}
	res2 := calculateSum(numbers...)
	fmt.Println("Sum of unpacked slice:", res2)
}
