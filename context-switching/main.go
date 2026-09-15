package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// TaskState simulates a lightweight Process Control Block (PCB) structure
type TaskState struct {
	TaskID    int
	Completed int
	Total     int
}

// workerTask simulates a long-running process that periodically yields execution control,
// triggering a Go runtime Goroutine Context Switch.
func workerTask(id int, totalSteps int, wg *sync.WaitGroup, stateChan chan<- TaskState) {
	defer wg.Done()

	for step := 1; step <= totalSteps; step++ {
		// Report execution state before yielding (similar to updating PCB state)
		stateChan <- TaskState{
			TaskID:    id,
			Completed: step,
			Total:     totalSteps,
		}

		// Artificial work simulation
		time.Sleep(10 * time.Millisecond)

		// Explicitly yield CPU time slice back to the Go Scheduler (User-space Context Switch)
		// This causes the Go M:N scheduler to pause this Goroutine and swap in another.
		runtime.Gosched()
	}
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 1. GOROUTINE CONTEXT SWITCHING & STATE MONITOR ")
	fmt.Println("==================================================")

	// Display available logical CPU cores used by the Go Runtime Scheduler
	fmt.Printf("Active OS Threads Available (GOMAXPROCS): %d\n\n", runtime.NumCPU())

	var wg sync.WaitGroup
	stateChan := make(chan TaskState, 20)

	numWorkers := 3
	stepsPerWorker := 4

	// Launch concurrent Goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workerTask(i, stepsPerWorker, &wg, stateChan)
	}

	// Goroutine to close state channel after all worker tasks finish
	go func() {
		wg.Wait()
		close(stateChan)
	}()

	// Monitor worker task states as context switches occur interleavingly
	fmt.Println("Execution Log (Interleaved Execution via Runtime Scheduler):")
	for state := range stateChan {
		fmt.Printf("[Task %d] Execution Step: %d/%d | Goroutines Active: %d\n",
			state.TaskID, state.Completed, state.Total, runtime.NumGoroutine())
	}

	fmt.Println("\nAll concurrent worker tasks completed successfully.")
}
