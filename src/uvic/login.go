package uvic

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/html"
)

func CreateAuthenticatedClient(username string, password string) (*http.Client, bool) {
	LOGIN_URL := "https://www.uvic.ca/cas/login?service=https%3A%2F%2Fwww.uvic.ca%2Ftools%2Findex.php"

	//create cookie jar to save cookies between requests
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

	//find our execution string from the login page html
	execution, found := findExecutionString(resp.Body)
	if !found {
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

	//uvic may return ok status code even with failed login this way we can check if we entered the protected page
	if resp.Request.URL.String() != "https://www.uvic.ca/tools/index.php" {
		return nil, false
	}

	return client, true
}

// piece of data required for login that is found on the login page html
// this function returns that string and if it was found
func findExecutionString(data io.Reader) (string, bool) {

	tokenizer := html.NewTokenizer(data)
	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			return "", false
		}

		if tokenType != html.SelfClosingTagToken {
			continue
		}

		tagName, hasAttr := tokenizer.TagName()
		if string(tagName) != "input" || hasAttr == false {
			continue
		}

		executionString := ""
		found := false
		for {
			keyBytes, valBytes, more := tokenizer.TagAttr()
			key := string(keyBytes)
			val := string(valBytes)
			if key == "value" {
				executionString = val
			} else if key == "name" {
				if val != "execution" {
					found = false
					break
				}
				found = true
			}
			if !more {
				break
			}

		}
		if !found {
			continue
		}
		return executionString, true
	}
	return "", false
}
