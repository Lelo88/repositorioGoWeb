package main

import (
	"fmt"
	"net/http"
)
// IndexHandler handles requests to the root URL path. It responds with a simple "Hello, World!" message.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}


func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", IndexHandler)

	// Start the server on port 8080
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)

	// ListenAndServe blocks the main goroutine, so it keeps the server running
	http.ListenAndServe(port, router)
}
