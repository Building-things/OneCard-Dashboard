package main

import (
	uvicapi "ONECard-Dashboard/uvic-api"
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

	fmt.Println("Server Running On Port 5000")
	fmt.Println("	http://localhost:5000/")
	http.ListenAndServe(":5000", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	log(r)
	loginTemplate.Execute(w, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	log(r)

	renderData := uvicapi.ONECardData{}

	if r.Method == "POST" {
		//populate form data with form values from request
		err := r.ParseForm()
		if err != nil {
			http.Redirect(w, r, "login/?err=formParseError", 302)
			return
		}

		//make sure the posted data has the correct format
		if !r.Form.Has("username") || !r.Form.Has("password") {
			http.Redirect(w, r, "login/?err=formWrongKeys", 302)
			return
		}

		//create an athenticated uvic client
		client, ok := uvicapi.CreateAuthenticatedClient(r.Form.Get("username"), r.Form.Get("password"))
		if !ok {
			http.Redirect(w, r, "login/?err=invalidCredentials", 302)
			return
		}

		//populate our render data
		renderData = uvicapi.GetOneCardData(client)
	}

	homeTemplate.Execute(w, renderData)
}

func log(r *http.Request) {
	fmt.Printf("[%s] %s %s\n", r.Host, r.Method, r.URL)
}
