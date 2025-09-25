package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// IndexHandler handles requests to the root URL path. It responds with a simple "Hello, World!" message.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}

// NewGame handles requests to the /new-game URL path. It responds with a "New Game!" message.
func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New Game!")
}

// LoadGame handles requests to the /load-game URL path. It responds with a "Load Game!" message.
func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Game!")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Play!")
}

// SaveGame handles requests to the /save-game URL path. It responds with a "Save Game!" message.
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About!")
}