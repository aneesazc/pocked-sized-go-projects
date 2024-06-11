package main

import (
	"html/template"
	"log"
	"net/http"
)

func write(w http.ResponseWriter, msg string){
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}

}

func renderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	errorCheck(err)
	err = parsedTemplate.Execute(w, nil)
	errorCheck(err)
}

func errorCheck(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, request *http.Request){
	renderTemplate(w, "home.page.tmpl")
}

func main() {
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}