// main.go - Hjertet av appen
// Her ligger:
// 1. Routes    - hvilken URL viser hvilken side (http.HandleFunc)
// 2. Static    - serverer CSS og andre filer fra /static/ mappen
// 3. Server    - start serveren (http.ListenAndServe) - skal alltid være SIST

package main

import (
	"ff-htmx-go/internal/components" // Bruk navnet fra go.mod + mappenavnet
	"html/template"
	"net/http"
)

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))
	http.Handle("/components/", http.StripPrefix("/components/", http.FileServer(http.Dir("/components"))))
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("internal/templates/index.html"))
        tmpl.Execute(w, nil)
    })

	http.HandleFunc("/test-side", func(w http.ResponseWriter, r *http.Request) {

        // 1️⃣ Først lager vi dataene
        data := components.CardData{
	        ID: "123",
	        Title: "Hei fra Go",
	        Description: "Dette er kortet, og du kan nå GO lang",
        } 

        // 2️⃣ Så genererer vi HTML-fragmentet
        // Vi henter HTML-strengen
        htmlFragment := components.RenderCard(data)

        // 3️⃣ Så laster vi templaten
        tmpl := template.Must(template.ParseFiles("internal/templates/test-side.html"))


        // 4️⃣ Vi sender fragmentet inn i templaten
		w.Write([]byte(htmlFragment))
		tmpl.Execute(w, nil)
    //     tmpl.Execute(w, map[string]any{
	//         "Card": template.HTML(htmlFragment),
    // })
		// Vi lager dataene

		// Her må vi sende htmlFragment inn i templaten vår eller skrive direkte
	})


// Start serveren på port 8080 - skal alltid være sist i main()
    http.ListenAndServe(":8080", nil)
}

