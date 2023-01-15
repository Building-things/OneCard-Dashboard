package main

import (
	uvicapi "ONECard-Dashboard/uvic"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const ONECARD_URL string = "https://www.uvic.ca/MyCard/account/summary"

func main() {

	//setup static file server
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//setup redirect to go to login page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log(r)
		http.Redirect(w, r, "/static/login.html", http.StatusPermanentRedirect)
	})

	http.HandleFunc("/home", homeHandler)

	//start the server
	fmt.Println("Server Running On Port 5000")
	fmt.Println("	http://localhost:5000/")
	http.ListenAndServe("0.0.0.0:5000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log(r)
	testing := true
	if r.Method == "POST" && !testing {
		//populate form data with form values from request
		err := r.ParseForm()
		if err != nil {
			http.Redirect(w, r, "/static/login.html/?err=formParseError", http.StatusFound)
			return
		}

		//make sure the posted data has the correct format
		if !r.Form.Has("username") || !r.Form.Has("password") {
			http.Redirect(w, r, "/static/login.html/?err=formWrongKeys", http.StatusFound)
			return
		}

		isBetaUser := false
		for _, val := range betaUsers {
			if r.Form.Get("username") == val {
				isBetaUser = true
				break
			}
		}

		if !isBetaUser {
			http.Redirect(w, r, "/static/login.html/?err=nonBetaUser", http.StatusFound)
			return
		}

		//create an athenticated uvic client
		client, ok := uvicapi.CreateAuthenticatedClient(r.Form.Get("username"), r.Form.Get("password"))
		if !ok {
			http.Redirect(w, r, "/static/login.html/?err=invalidCredentials", http.StatusFound)
			return
		}
		//get the onecard page html
		resp, err := client.Get(ONECARD_URL)
		if err != nil {
			http.Redirect(w, r, "/static/login.html/?err=UvicServerError", http.StatusFound)
			return
		}

		//convert the response to bytes
		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	http.Redirect(w, r, "/static/login.html/?err=MalformedResponseBytes", http.StatusFound)
		// 	return
		// }

		parser, err := html.Parse(resp.Body)
		if err != nil {
			http.Redirect(w, r, "/static/login.html/?err=MalformedResponseBytes", http.StatusFound)
			return
		}

		//inject our script tag into the response html so we can take over on the client side
		injected_script_tag_string := "<script defer src=\"/static/js/app.js\">"
		reader := strings.NewReader(injected_script_tag_string)
		parser_2, err := html.Parse(reader)
		if err != nil {
			http.Redirect(w, r, "/static/login.html/?err=MalformedPayloadBytes", http.StatusFound)
			return
		}

		//add our script tag to the original response
		parser.AppendChild(parser_2)

		//render the new html
		html.Render(w, parser)
	} else {
		http.Redirect(w, r, "/static/test.html", http.StatusPermanentRedirect)
	}
}

func log(r *http.Request) {
	fmt.Printf("[%s] %s %s\n", r.Host, r.Method, r.URL)
}
