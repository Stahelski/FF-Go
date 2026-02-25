// main.go - Hjertet av appen
// Her ligger:
// 1. Routes    - hvilken URL viser hvilken side (http.HandleFunc)
// 2. Static    - serverer CSS og andre filer fra /static/ mappen
// 3. Server    - start serveren (http.ListenAndServe) - skal alltid være SIST

package main

import (
	"html/template"
	"net/http"
)

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/index.html"))
        tmpl.Execute(w, nil)
    })

	http.HandleFunc("/test-side", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/test-side.html"))
		tmpl.Execute(w, nil)
	})
// Start serveren på port 8080 - skal alltid være sist i main()
    http.ListenAndServe(":8080", nil)
}

