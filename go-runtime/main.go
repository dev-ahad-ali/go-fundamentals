package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

// printRuntimeStats retrieves and displays current metrics managed by the Go runtime
func printRuntimeStats(label string) {
	var memStats runtime.MemStats
	// Read current memory allocation statistics from the Go runtime
	runtime.ReadMemStats(&memStats)

	fmt.Printf("--- Runtime Stats [%s] ---\n", label)
	fmt.Printf("Logical CPUs (GOMAXPROCS): %d\n", runtime.NumCPU())
	fmt.Printf("Active Goroutines:        %d\n", runtime.NumGoroutine())
	fmt.Printf("Allocated Heap Memory:   %d KB\n", memStats.HeapAlloc/1024)
	fmt.Printf("Total Garbage Collections: %d\n", memStats.NumGC)
	fmt.Println("-----------------------------------")
}

// homeHandler processes requests to the root endpoint
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Log the address handling this request
	log.Printf("Handling request for / from %s on Goroutine ID context", r.RemoteAddr)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Go Runtime HTTP Server!"))
}

// aboutHandler processes requests to the /about endpoint
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling request for /about from %s", r.RemoteAddr)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("About Page: Powered by Go embedded runtime and Netpoller."))
}

func main() {
	// 1. Display initial runtime status on startup
	printRuntimeStats("Startup")

	// 2. Initialize HTTP Router (ServeMux)
	mux := http.NewServeMux()

	// Register route paths to custom handler functions
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	// 3. Trigger manual Garbage Collection to observe runtime behavior
	runtime.GC()
	printRuntimeStats("After Manual GC")

	serverPort := ":8080"
	fmt.Printf("Starting HTTP server on port %s...\n", serverPort)

	// 4. Configure HTTP Server instance
	server := &http.Server{
		Addr:         serverPort,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 5. Start listening for incoming connections
	// The Go Netpoller accepts sockets and assigns each connection to a new Goroutine
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server stopped with error: %v", err)
	}
}
