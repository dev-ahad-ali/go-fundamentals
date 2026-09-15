package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Counter is allocated on the shared Heap memory space.
// All goroutines can read/write to this structure.
type Counter struct {
	mu    sync.Mutex
	value int
}

// worker represents a function executed by an independent Goroutine.
// Each goroutine maintains its own call stack frame.
func worker(id int, sharedCounter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()

	// localStackVar resides entirely on this specific Goroutine's isolated stack frame.
	localStackVar := id * 100

	// Simulate work execution
	time.Sleep(10 * time.Millisecond)

	// Accessing shared Heap memory safely using Mutex locking
	sharedCounter.mu.Lock()
	sharedCounter.value++
	currentCounterVal := sharedCounter.value
	sharedCounter.mu.Unlock()

	// Print isolated stack values vs shared heap values
	fmt.Printf("Goroutine Worker ID: %d | Isolated Stack Variable: %d | Shared Heap Counter: %d\n",
		id, localStackVar, currentCounterVal)
}

func main() {
	var wg sync.WaitGroup

	// Allocate shared structure on the heap
	sharedCounter := &Counter{}

	totalGoroutines := 5

	fmt.Printf("Total Available Logical CPU Cores: %d\n", runtime.NumCPU())
	fmt.Printf("Launching %d Goroutines (each with independent initial ~2KB stacks)...\n\n", totalGoroutines)

	// Launching goroutines concurrently
	for i := 1; i <= totalGoroutines; i++ {
		wg.Add(1)
		go worker(i, sharedCounter, &wg)
	}

	// Wait for all goroutines to complete execution
	wg.Wait()

	fmt.Printf("\nExecution Completed. Final Shared Heap Counter Value: %d\n", sharedCounter.value)
}
