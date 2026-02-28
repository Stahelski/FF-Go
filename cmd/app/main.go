// main.go - Hjertet av appen
// Her ligger:
// 1. Routes    - hvilken URL viser hvilken side (http.HandleFunc)
// 2. Static    - serverer CSS og andre filer fra /static/ mappen
// 3. Server    - start serveren (http.ListenAndServe) - skal alltid være SIST

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Parse ALLE templates én gang ved oppstart
var templates = template.Must(
	template.ParseGlob("internal/templates/*.html"),
)

func main() {

	// Routes
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/projects", projects)

	// Static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ---------------- ROUTES ----------------

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, "home.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func projects(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html")
}

// ---------------- TEMPLATE RENDER ----------------

func renderTemplate(w http.ResponseWriter, name string) {
	err := templates.ExecuteTemplate(w, name, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}