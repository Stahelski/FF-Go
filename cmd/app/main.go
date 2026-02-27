// main.go - Hjertet av appen
// Her ligger:
// 1. Routes    - hvilken URL viser hvilken side (http.HandleFunc)
// 2. Static    - serverer CSS og andre filer fra /static/ mappen
// 3. Server    - start serveren (http.ListenAndServe) - skal alltid være SIST

//! 14:24
// https://www.youtube.com/watch?v=o90ZnlorYwA

package main

import (
	"html/template"
	"log"
	"net/http"
)

func main(){
	// Route handler for home page (root URL)
http.HandleFunc("/", home)
http.HandleFunc("/projects", projects)
http.HandleFunc("/about", about)

// Start server
err := http.ListenAndServe(":8080", nil)
if err != nil{
	log.Fatal(err)
}
}

//Home func handels requests to home
func home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "home.html")
}

func projects(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "internal/templates/about.html") //! Funker denne route?
}

func about(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "/internal/templates/about.html") //! Funker denne route?
}



func renderTemplate(w http.ResponseWriter, tmpl string){
	// Parsing the specified template file being passed as input (home, about, etc...)
	t, err := template.ParseFiles("/internal/templates" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}

