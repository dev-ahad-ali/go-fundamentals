package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Task represents a process entry for scheduling simulation
type Task struct {
	ID        string
	BurstTime int // Remaining CPU time units needed
}

// simulateCPUScheduler demonstrates a simple Round Robin (RR) time-slicing scheduler algorithm
func simulateCPUScheduler(tasks []Task, timeQuantum int) {
	fmt.Println("--------------------------------------------------")
	fmt.Println(" SIMULATING ROUND-ROBIN (RR) CPU SCHEDULER")
	fmt.Println("--------------------------------------------------")

	queue := make([]Task, len(tasks))
	copy(queue, tasks)

	timeElapsed := 0

	for len(queue) > 0 {
		currentTask := queue[0]
		queue = queue[1:] // Dequeue

		execTime := timeQuantum
		if currentTask.BurstTime < timeQuantum {
			execTime = currentTask.BurstTime
		}

		currentTask.BurstTime -= execTime
		timeElapsed += execTime

		fmt.Printf("[Time: %02dms] Executed Task %s for %dms | Remaining: %dms\n",
			timeElapsed, currentTask.ID, execTime, currentTask.BurstTime)

		if currentTask.BurstTime > 0 {
			queue = append(queue, currentTask) // Re-queue preempted task
		} else {
			fmt.Printf("   ==> Task %s COMPLETED at %dms\n", currentTask.ID, timeElapsed)
		}
	}
}

// worker performs intensive computational work to demonstrate parallelism across CPU cores
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()

	// CPU-bound loop simulation
	count := 0
	for i := 0; i < 50_000_000; i++ {
		count += i
	}

	fmt.Printf("Worker %d finished on OS Thread | Duration: %v\n", id, time.Since(start))
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" 1. CONCURRENCY VS PARALLELISM DEMO")
	fmt.Println("==================================================")

	numCPU := runtime.NumCPU()
	fmt.Printf("Available Physical/Logical CPU Cores: %d\n", numCPU)

	// Scenario A: Single Core Execution (Concurrency via Time-Slicing)
	fmt.Println("\n--- Scenario A: Concurrency (GOMAXPROCS = 1) ---")
	runtime.GOMAXPROCS(1)

	var wg1 sync.WaitGroup
	startSingle := time.Now()
	for i := 1; i <= 3; i++ {
		wg1.Add(1)
		go worker(i, &wg1)
	}
	wg1.Wait()
	fmt.Printf("Total Time (Single Core / Concurrency): %v\n", time.Since(startSingle))

	// Scenario B: Multi-Core Execution (True Parallelism)
	fmt.Println("\n--- Scenario B: Parallelism (GOMAXPROCS = Max Cores) ---")
	runtime.GOMAXPROCS(numCPU)

	var wg2 sync.WaitGroup
	startMulti := time.Now()
	for i := 1; i <= 3; i++ {
		wg2.Add(1)
		go worker(i, &wg2)
	}
	wg2.Wait()
	fmt.Printf("Total Time (Multi Core / Parallelism) : %v\n", time.Since(startMulti))

	// Section 2: OS CPU Scheduling Algorithm Simulation
	fmt.Println("\n==================================================")
	fmt.Println(" 2. CPU SCHEDULING ALGORITHM SIMULATION")
	fmt.Println("==================================================")

	taskList := []Task{
		{ID: "Process-A", BurstTime: 10},
		{ID: "Process-B", BurstTime: 4},
		{ID: "Process-C", BurstTime: 6},
	}

	// Quantum = 3ms
	simulateCPUScheduler(taskList, 3)
}
