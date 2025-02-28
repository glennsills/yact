package ui

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles(
	"ui/views/default.html",
	"ui/views/users/listusers.html"))

var validPath = regexp.MustCompile("^/(edit|save|view|users)/([a-zA-Z0-9]+)$")

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "default.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "listusers.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func assignHandlers() {
	// Log the names of all parsed templates
	log.Println("Parsed templates:", templates.DefinedTemplates())

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/users/", listHandler)
}
