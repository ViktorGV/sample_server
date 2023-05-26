package main

import (
	"fmt"
	"log"
	"net"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read request body: %v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Get the IP address of the caller
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("Failed to get caller's IP address: %v", err)
	}

	// Print the request body and caller's IP address
	log.Printf("Request Body: %s\n", body)
	log.Printf("Caller IP: %s\n", ip)

	// Optionally, you can respond to the request with a custom message
	message := "Request received successfully"
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
