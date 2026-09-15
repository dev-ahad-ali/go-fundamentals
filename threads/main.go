package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Shared Counter Structure representing shared memory among threads
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment safely updates the shared memory using a Mutex lock
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value safely reads the counter value
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// workerTask executes as an independent Goroutine (M:N User Thread)
func workerTask(id int, counter *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[Worker %d] Started execution on OS Thread | Active Goroutines: %d\n",
		id, runtime.NumGoroutine())

	for i := 0; i < 5; i++ {
		counter.Increment()
		// Sleep causes the Go scheduler to pause this goroutine
		// and perform a lightweight user-space context switch.
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Printf("[Worker %d] Completed its task.\n", id)
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" DEMO: MULTITHREADING & GOROUTINE SCHEDULING")
	fmt.Println("==================================================")

	// Display OS / Machine thread capabilities
	fmt.Printf("Physical / Logical CPU Cores available: %d\n", runtime.NumCPU())
	fmt.Printf("Default Max OS Threads (GOMAXPROCS): %d\n\n", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup
	counter := &SafeCounter{}

	numWorkers := 4

	// Launch multiple Goroutines (Multiplexed on OS Kernel Threads)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workerTask(i, counter, &wg)
	}

	// Wait for all worker threads/goroutines to finish execution
	wg.Wait()

	fmt.Println("\n--------------------------------------------------")
	fmt.Printf("Final Counter Value (Shared Memory Updated): %d\n", counter.Value())
	fmt.Println("All Goroutines finished execution cleanly.")
}
