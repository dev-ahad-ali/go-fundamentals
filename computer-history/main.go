package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// User defines a struct to demonstrate memory layout & alignment in RAM
type User struct {
	ID       uint64 // 8 bytes
	IsActive bool   // 1 byte
	Age      uint8  // 1 byte
	// Note: 6 bytes of padding inserted by compiler on 64-bit systems for alignment!
	Score float64 // 8 bytes
}

// Function demonstrating stack vs heap memory escape
func createStackUser() User {
	// Allocated on function stack frame
	u := User{ID: 101, IsActive: true, Age: 25, Score: 98.5}
	return u // Returned by value (copied)
}

func createHeapUser() *User {
	// Escapes to heap because a pointer is returned
	u := User{ID: 102, IsActive: true, Age: 30, Score: 88.0}
	return &u
}

func main() {
	fmt.Println("==================================================")
	fmt.Println("   GO ARCHITECTURE & MEMORY INSPECTION PROGRAM   ")
	fmt.Println("==================================================")

	// 1. Inspect System Architecture & OS Environment
	fmt.Printf("Operating System:     %s\n", runtime.GOOS)
	fmt.Printf("CPU Architecture:     %s\n", runtime.GOARCH)
	fmt.Printf("Logical CPU Cores:    %d\n", runtime.NumCPU())

	// Word size calculation (32-bit vs 64-bit)
	wordBits := 8 * unsafe.Sizeof(uintptr(0))
	fmt.Printf("System Word Size:     %d-bit\n", wordBits)

	fmt.Println("\n--------------------------------------------------")
	fmt.Println(" 2. Memory Units & Data Type Sizes in RAM (Bytes)")
	fmt.Println("--------------------------------------------------")

	var b bool
	var i8 int8
	var i32 int32
	var i64 int64
	var f64 float64
	var str string
	var ptr *int

	fmt.Printf("bool size:            %d Byte  (%d bit)\n", unsafe.Sizeof(b), unsafe.Sizeof(b)*8)
	fmt.Printf("int8 size:            %d Byte  (%d bits)\n", unsafe.Sizeof(i8), unsafe.Sizeof(i8)*8)
	fmt.Printf("int32 size:           %d Bytes (%d bits)\n", unsafe.Sizeof(i32), unsafe.Sizeof(i32)*8)
	fmt.Printf("int64 size:           %d Bytes (%d bits)\n", unsafe.Sizeof(i64), unsafe.Sizeof(i64)*8)
	fmt.Printf("float64 size:         %d Bytes (%d bits)\n", unsafe.Sizeof(f64), unsafe.Sizeof(f64)*8)
	fmt.Printf("string header size:   %d Bytes (Pointer + Len)\n", unsafe.Sizeof(str))
	fmt.Printf("Pointer size:         %d Bytes (%d-bit memory address)\n", unsafe.Sizeof(ptr), wordBits)

	fmt.Println("\n--------------------------------------------------")
	fmt.Println(" 3. Struct Memory Layout & Memory Alignment")
	fmt.Println("--------------------------------------------------")

	u := User{ID: 1, IsActive: true, Age: 28, Score: 95.4}
	fmt.Printf("User Struct Total Size: %d Bytes\n", unsafe.Sizeof(u))

	// Inspecting memory offsets of individual fields inside the struct
	fmt.Printf("  - Offset of ID:       %d Bytes\n", unsafe.Offsetof(u.ID))
	fmt.Printf("  - Offset of IsActive: %d Bytes\n", unsafe.Offsetof(u.IsActive))
	fmt.Printf("  - Offset of Age:      %d Bytes\n", unsafe.Offsetof(u.Age))
	fmt.Printf("  - Offset of Score:    %d Bytes (aligned to 8-byte boundary)\n", unsafe.Offsetof(u.Score))

	fmt.Println("\n--------------------------------------------------")
	fmt.Println(" 4. RAM Address Inspection (Stack vs. Heap)")
	fmt.Println("--------------------------------------------------")

	stackVal := createStackUser()
	heapVal := createHeapUser()

	fmt.Printf("Stack Object Value:   %+v\n", stackVal)
	fmt.Printf("Stack Object Address: %p\n", &stackVal)

	fmt.Printf("Heap Object Value:    %+v\n", heapVal)
	fmt.Printf("Heap Object Address:  %p (allocated dynamically in RAM heap)\n", heapVal)

	fmt.Println("\n--------------------------------------------------")
	fmt.Println(" 5. Bitwise Operations (Low-Level Hardware Flags)")
	fmt.Println("--------------------------------------------------")

	// Bitwise flags representing system permissions
	const (
		ReadPermission  = 1 << 0 // 0001 (1)
		WritePermission = 1 << 1 // 0010 (2)
		ExecPermission  = 1 << 2 // 0100 (4)
	)

	// Combine Read and Write permissions using bitwise OR (|)
	userPerm := ReadPermission | WritePermission // 0011 (3)

	fmt.Printf("Permission Mask Value: %d (Binary: %04b)\n", userPerm, userPerm)
	fmt.Printf("Has Read Permission?   %t\n", (userPerm&ReadPermission) != 0)
	fmt.Printf("Has Exec Permission?   %t\n", (userPerm&ExecPermission) != 0)
}
