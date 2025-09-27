package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir  = "templates/"
	baseTemplate = templateDir + "base.html"
)

// IndexHandler handles requests to the root URL path. It responds with a simple "Hello, World!" message.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, baseTemplate, "index.html", nil)
}

// NewGame handles requests to the /new-game URL path. It responds with a "New Game!" message.
func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, baseTemplate, "new-game.html", nil)
}

// Game handles requests to the /game URL path. It responds with a "Game!" message.
func Game(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, baseTemplate, "game.html", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Play!")
}

// SaveGame handles requests to the /save-game URL path. It responds with a "Save Game!" message.
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, baseTemplate, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, base, page string, data any) {
	tmpl := template.Must(template.ParseFiles(base, templateDir + page))

	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}