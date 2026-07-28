package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)


type Fruit struct {
	Name string
	Description string
}

func getFruits(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/fruits.html", "templates/fruit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	type PageData struct {
		Fruits []Fruit
	}

	data := PageData{
		Fruits: []Fruit{
			{Name: "Apple", Description: "Not the company"},
			{Name: "Banana", Description: "A long yellow fruit"},
		},
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println("Execution error:", err)
	}
}

func postFruits(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1 MB limit

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	data := Fruit{
		Name: r.PostForm.Get("name"), 
		Description: r.PostForm.Get("description"),
	}

	tmpl, err := template.ParseFiles("templates/fruit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	err = tmpl.ExecuteTemplate(w, "fruit", data)
	if err != nil {
		log.Println("Execution error:", err)
	}
}

func deleteFruit(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
	fmt.Fprint(w, "")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /fruits", postFruits)
	mux.HandleFunc("GET /fruits", getFruits)
	mux.HandleFunc("DELETE /fruits/{id}", deleteFruit) 

	fileServer := http.FileServer(http.Dir("./static"))
	staticHandler := http.StripPrefix("/static/", fileServer)
	mux.Handle("GET /static/", staticHandler)

	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}