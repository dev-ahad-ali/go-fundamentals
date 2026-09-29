package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// helloHandler handles requests directed to the "/hello" route path.
// w (http.ResponseWriter) allows sending response headers and body to the OS socket.
// r (*http.Request) contains parsed request details received from the network file descriptor.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Log details retrieved from the parsed HTTP request context
	log.Printf("Received %s request for %s from client address %s", r.Method, r.URL.Path, r.RemoteAddr)

	// Set HTTP response headers
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Write response payload directly back to the HTTP connection stream
	responseText := fmt.Sprintf("Hello from Go HTTP Server! Server Time: %s", time.Now().Format(time.RFC3339))
	_, err := w.Write([]byte(responseText))
	if err != nil {
		log.Printf("Failed writing response to socket buffer: %v", err)
	}
}

// statusHandler returns a simple JSON response indicating server health.
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "HEALTHY", "runtime": "Go net/http"}`))
}

func main() {
	// 1. Instantiate a new HTTP Request Multiplexer (ServeMux)
	// ServeMux handles URL pattern matching and routes requests to appropriate handlers.
	mux := http.NewServeMux()

	// 2. Register route handlers with the multiplexer
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/status", statusHandler)

	serverAddress := ":8080"
	log.Printf("Starting Go HTTP Server on port %s...", serverAddress)

	// 3. Configure HTTP server timeouts and parameters
	// Timeout configurations prevent resource exhaustion from slow or stale socket connections.
	server := &http.Server{
		Addr:         serverAddress,
		Handler:      mux,              // Custom router multiplexer
		ReadTimeout:  5 * time.Second,  // Max duration for reading full request from OS socket
		WriteTimeout: 10 * time.Second, // Max duration for writing response back to OS socket
		IdleTimeout:  15 * time.Second, // Max duration to keep keep-alive TCP connections open
	}

	// 4. Start listening on the TCP port
	// ListenAndServe binds to the network port, listens for connections via OS kernel system calls,
	// and passes incoming connection file descriptors to the Go runtime Netpoller.
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP Server stopped unexpectedly: %v", err)
	}
}
