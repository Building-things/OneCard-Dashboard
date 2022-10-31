package onecardparsing

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func createAuthenticatedClient(username string, password string) (*http.Client, bool) {
	LOGIN_URL := "https://www.uvic.ca/cas/login?service=https%3A%2F%2Fwww.uvic.ca%2Ftools%2Findex.php"

	//create cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, false
	}

	//make our client for the http requests
	client := &http.Client{
		Jar: jar,
	}

	//get the login url to set up some cookies
	resp, err := client.Get(LOGIN_URL)
	if err != nil {
		return nil, false
	}

	//find our execution string to send with the username and password
	doc, doc_err := goquery.NewDocumentFromReader(resp.Body)
	if doc_err != nil {
		return nil, false
	}
	execution, exists := doc.Find("input[name='execution']").Attr("value")
	if !exists {
		return nil, false
	}

	//send the login request
	resp, err = client.PostForm(LOGIN_URL,
		url.Values{
			"username":  {username},
			"password":  {password},
			"execution": {execution},
			"_eventId":  {"submit"},
		},
	)
	if err != nil {
		return nil, false
	}

	//if we are not here we did not succesfully login
	if resp.Request.URL.String() != "https://www.uvic.ca/tools/index.php" {
		return nil, false
	}
	return client, true
}
