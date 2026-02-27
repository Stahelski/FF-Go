
package components // Dette må stå øverst i alle filer i denne mappen

import (
	"bytes"
	"html/template"
)

// CardData må starte med stor bokstav for å kunne brukes i main.go. STOR bokstav = EXPORT (synlig i main.go)
type CardData struct {
	ID          string
	Title       string
	Description string
}

// RenderCard må starte med stor bokstav for å være tilgjengelig utenfor denne mappen
func RenderCard(data CardData) string {
	tmpl := `
	<div id="card-{{.ID}}" class="card">
		<h3>{{.Title}}</h3>
		<p>{{.Description}}</p>
		<button 
			data-on-click="$$get('/api/card/{{.ID}}')"
			class="card-button">
			Card button
		</button>
	</div>
	`
	t := template.Must(template.New("card").Parse(tmpl))
	var tpl bytes.Buffer
	t.Execute(&tpl, data)

	return tpl.String()
}