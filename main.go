package main

import (
	"log"
	"net/http"

	"github.com/Lelo88/repositorioGoWeb/handlers"
)


func main() {

	// Create a new ServeMux (router)
	router := http.NewServeMux()

	// Serve static files from the "static" directory
	fileServer := http.FileServer(http.Dir("./static/"))
	// Rute to serve static files when the URL path starts with /static/
	router.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Register the IndexHandler for the root URL path
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/new-game", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	// Start the server on port 8080
	port := ":8080"
	log.Printf("Server is running on http://localhost%s\n", port)

	// ListenAndServe blocks the main goroutine, so it keeps the server running
	http.ListenAndServe(port, router)
}
