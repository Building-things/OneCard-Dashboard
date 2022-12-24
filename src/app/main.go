package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// set up the templates here we only have two pages so might aswell make them global
// to avoid passing between functions.
var loginTemplate = template.Must(template.ParseFiles("../templates/login.html"))
var homeTemplate = template.Must(template.ParseFiles("../templates/home.html"))

func main() {

	//set up file server for css, js and images
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", login)
	http.HandleFunc("/home", home)

	http.ListenAndServe(":5000", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	log(r)
	switch r.Method {
	case "GET":
		loginTemplate.Execute(w, nil)
	case "POST":
		w.Write([]byte("...pretending to login"))
		//create object with our home template data from login
		http.Redirect(w, r, "/home", 302)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	log(r)
	homeTemplate.Execute(w, nil)
}

func log(r *http.Request) {
	fmt.Printf("[%s] %s %s\n", r.Host, r.Method, r.URL)
}
