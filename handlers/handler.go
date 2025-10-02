package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Player struct {
	Name string
}

var player Player

const (
	templateDir  = "templates/"
	baseTemplate = templateDir + "base.html"
)

// IndexHandler handles requests to the root URL path. It responds with a simple "Hello, World!" message.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Restart player values
	restartValues()
	// Render the index template
	renderTemplate(w, baseTemplate, "index.html", nil)
}

// NewGame handles requests to the /new-game URL path. It responds with a "New Game!" message.
func NewGame(w http.ResponseWriter, r *http.Request) {
	// Restart player values
	restartValues()
	renderTemplate(w, baseTemplate, "new-game.html", nil)
}

// Game handles requests to the /game URL path. It responds with a "Game!" message.
func Game(w http.ResponseWriter, r *http.Request) {
	// Handle form submission
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		
		// Get the player's name from the form data
		player.Name = r.Form.Get("name")
	}

	if player.Name == "" {
		http.Redirect(w, r, "/new-game", http.StatusFound)
		return
	}
	//fmt.Println("Player Name:", player.Name)

	renderTemplate(w, baseTemplate, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Play!")
}

// SaveGame handles requests to the /save-game URL path. It responds with a "Save Game!" message.
func About(w http.ResponseWriter, r *http.Request) {
	// Restart player values
	restartValues()
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

// Reinit values
func restartValues() {
	player.Name = ""
}