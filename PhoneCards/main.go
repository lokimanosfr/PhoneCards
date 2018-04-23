package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	ID    string
	Title string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	fmt.Println("Listening on port: 3000")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/card-view", cardView)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("./src/"))))
	http.ListenAndServe(":3000", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//t, err := template.ParseFiles("./templates/index.gohtml", "./templates/header.gohtml", "./templates/footer.gohtml", "./templates/card.gohtml", "./templates/jumbotron.gohtml")
	err := tpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}

}

func cardView(w http.ResponseWriter, req *http.Request) {
	pd := pageData{
		Title: "ID = 1",
		ID:    req.FormValue("id"),
	}

	//t, err := template.ParseFiles("./templates/card-view.gohtml", "./templates/header.gohtml", "./templates/footer.gohtml")
	err := tpl.ExecuteTemplate(w, "card-view", pd)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

}
