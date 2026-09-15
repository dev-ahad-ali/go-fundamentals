package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 1. PROCESS & OS ENVIRONMENT INSPECTION")
	fmt.Println("==================================================")

	// Fetch current Process ID (PID) assigned by the Operating System
	pid := os.Getpid()
	parentPid := os.Getppid()
	fmt.Printf("Current Process ID (PID) : %d\n", pid)
	fmt.Printf("Parent Process ID (PPID) : %d\n", parentPid)

	// Fetch OS Architecture and Platform
	fmt.Printf("Number of Logical CPUs   : %d\n", os.Getenv("NUMBER_OF_PROCESSORS"))

	fmt.Println("\n==================================================")
	fmt.Println(" 2. HIGH-LEVEL OS FILE CREATION & WRITE")
	fmt.Println("==================================================")

	fileName := "os_demo.txt"
	content := []byte("Hello from Go User Space to OS File System!\n")

	// Standard library os.WriteFile handles system calls under the hood
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Printf("Successfully wrote file '%s' via standard os package.\n", fileName)

	fmt.Println("\n==================================================")
	fmt.Println(" 3. DIRECT LOW-LEVEL SYSTEM CALL (syscall)")
	fmt.Println("==================================================")

	// Bypassing standard abstractions to perform raw System Calls
	// Opening file descriptor directly via syscall.Open
	// Flags: O_RDWR (Read/Write), O_APPEND (Append mode)
	fd, err := syscall.Open(fileName, syscall.O_RDWR|syscall.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Syscall Open Error:", err)
		return
	}

	// Defer closing the File Descriptor via Syscall
	defer syscall.Close(fd)

	fmt.Printf("Acquired OS File Descriptor (FD): %d\n", fd)

	// Writing directly to the file descriptor using syscall.Write
	appendData := []byte("Appended text via raw system call (syscall.Write).\n")
	n, err := syscall.Write(fd, appendData)
	if err != nil {
		fmt.Println("Syscall Write Error:", err)
		return
	}

	fmt.Printf("Wrote %d bytes directly using Kernel System Call.\n", n)

	// Clean up demo file
	_ = os.Remove(fileName)
	fmt.Println("\nCleaned up demo file.")
}
