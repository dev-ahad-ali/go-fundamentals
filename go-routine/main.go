package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// TaskResult holds the output data produced by a worker goroutine
type TaskResult struct {
	WorkerID int
	Result   int
	Duration time.Duration
}

// processTask represents a unit of work executed concurrently by a Goroutine
func processTask(id int, data int, resultsChan chan<- TaskResult, wg *sync.WaitGroup) {
	// Ensure the WaitGroup counter is decremented when the Goroutine completes
	defer wg.Done()

	startTime := time.Now()

	// Simulate computational or I/O work
	computation := data * data
	time.Sleep(50 * time.Millisecond)

	elapsed := time.Since(startTime)

	// Send the result back via the communication channel
	resultsChan <- TaskResult{
		WorkerID: id,
		Result:   computation,
		Duration: elapsed,
	}
}

func main() {
	// Print runtime environment configuration
	fmt.Printf("GOMAXPROCS (Logical Processors P): %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("Initial active Goroutines: %d\n\n", runtime.NumGoroutine())

	taskCount := 5
	var wg sync.WaitGroup

	// Buffered channel to collect results without blocking workers unnecessarily
	resultsChan := make(chan TaskResult, taskCount)

	// Launch multiple concurrent goroutines using the 'go' keyword
	for i := 1; i <= taskCount; i++ {
		wg.Add(1)
		go processTask(i, i*10, resultsChan, &wg)
	}

	fmt.Printf("Active Goroutines after spawning workers: %d\n", runtime.NumGoroutine())

	// Wait for all worker goroutines to complete in a separate monitoring goroutine
	go func() {
		wg.Wait()
		close(resultsChan) // Close channel once all workers are done
	}()

	// Read results from the channel as they become available
	fmt.Println("\nProcessing results received from channel:")
	for result := range resultsChan {
		fmt.Printf("Worker %d completed task. Input squared: %d (Took %v)\n",
			result.WorkerID, result.Result, result.Duration)
	}

	fmt.Printf("\nFinal active Goroutines: %d\n", runtime.NumGoroutine())
}
