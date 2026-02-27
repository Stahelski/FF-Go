// main.go - Hjertet av appen
// Her ligger:
// 1. Routes    - hvilken URL viser hvilken side (http.HandleFunc)
// 2. Static    - serverer CSS og andre filer fra /static/ mappen
// 3. Server    - start serveren (http.ListenAndServe) - skal alltid være SIST

package main

import (
	"ff-htmx-go/components" // Bruk navnet fra go.mod + mappenavnet
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
		// Vi lager dataene
		data := components.CardData{
			ID: "123",
			Title: "Hei fra Go",
			Description: "Dette er kortet, og du kan nå GO lang",
		}
		// Vi henter HTML-strengen
		htmlFragment := components.RenderCard(data)

		// Her må vi sende htmlFragment inn i templaten vår eller skrive direkte
		w.Write([]byte(htmlFragment))
	})


// Start serveren på port 8080 - skal alltid være sist i main()
    http.ListenAndServe(":8080", nil)
}

