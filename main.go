package main

import (
	"html/template"
	"log"
	"net/http"
)


type Fruit struct {
	Name string
	Description string
}

// PageData defines the structure of the data passed to the template
type PageData struct {
	Fruits []Fruit
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Parse the template file
	tmpl, err := template.ParseFiles("templates/index.html", "templates/fruit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Define the data to inject
	data := PageData{
		Fruits: []Fruit{
			{Name: "Apple", Description: "Not the company"},
			{Name: "Banana", Description: "A long yellow fruit"},
		},
	}

	// Render the template with data and write to response
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Execution error:", err)
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleHome)

	fileServer := http.FileServer(http.Dir("./static"))
	staticHandler := http.StripPrefix("/static/", fileServer)
	mux.Handle("GET /static/", staticHandler)

	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}