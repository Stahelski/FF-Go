Åpne appen [HUSK] du har Air innstalert

# Med "Air" installert så kjør dette i terminal for å starte prosjekt:
air

# Uten "Air" installert må du gjøre dette for å kjøre app:
Serveren starter ikke nettleseren automatisk (ulikt React/Vite).
Du må selv åpne nettleseren og gå til:
http://localhost:8080
Stopp serveren
Trykk Ctrl + C i terminalen.

Prosjektstruktur
FF-htmx-go/
├── main.go            # Hovedfil, router og serveroppsett
├── go.mod             # Go module-fil
├── static/
│   └── style.css      # CSS-stiler
├── templates/
│   └── index.html     # HTML-maler med Datastar
└── db/
    └── db.go          # Databaseoppsett (SQLite)

Forskjeller fra React
React / Vite                         Go + Datastar

npm run dev                          go run main.go
Åpner nettleser automatisk           Du skriver inn http://localhost:8080 selv
npm install                          go mod tidy
node_modules/                        Ingen, Go laster ned avhengigheter selv



# Ved feilmeldinger:
Skriv "go mod tidy" i terminal. 
rydder opp avhengigheter, tilsvarer npm install i React. Kjør den hvis du får feilmeldinger om manglende pakker.