package onecardparsing

import (
	"math"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Balances_T struct {
	StandardMealPlan float64
	PlusMealPlan     float64
	Flex             float64
}
type Transaction_T struct {
	Location string
	Amount   float64
	Date     string
}
type Transactions_T struct {
	Recent []Transaction_T
	All    []Transaction_T
}
type ONECardData struct {
	Balances     Balances_T
	Transactions Transactions_T
}

func OneCardData(username string, password string) (ONECardData, bool) {
	LOGIN_URL := "https://www.uvic.ca/cas/login?service=https%3A%2F%2Fwww.uvic.ca%2Ftools%2Findex.php"
	ONECARD_URL := "https://www.uvic.ca/MyCard/account/summary"

	//create cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		return ONECardData{}, false
	}

	//make our client for the http requests
	client := &http.Client{
		Jar: jar,
	}

	//get the login url to set up some cookies
	resp, err := client.Get(LOGIN_URL)
	if err != nil {
		return ONECardData{}, false
	}

	//find our execution string to send with the username and password
	doc, doc_err := goquery.NewDocumentFromReader(resp.Body)
	if doc_err != nil {
		return ONECardData{}, false
	}
	execution, exists := doc.Find("input[name='execution']").Attr("value")
	if !exists {
		return ONECardData{}, false
	}

	//send the login request
	_, err = client.PostForm(LOGIN_URL,
		url.Values{
			"username":  {username},
			"password":  {password},
			"execution": {execution},
			"_eventId":  {"submit"},
		},
	)
	if err != nil {
		return ONECardData{}, false
	}

	//get onecard page
	resp, err = client.Get(ONECARD_URL)
	if err != nil {
		return ONECardData{}, false
	}

	//parse user data from the onecard page
	doc, doc_err = goquery.NewDocumentFromReader(resp.Body)
	if doc_err != nil {
		return ONECardData{}, false
	}

	//since UVIC returns 200 even on failed login we have to see if we are still on the login page
	//if this exists we are on thus the login failed
	_, exists = doc.Find("input[name='execution']").Attr("value")
	if exists {
		return ONECardData{}, false
	}

	data := ONECardData{}
	//parse balances
	doc.Find(".summary.transactions > ul > li > .transaction-amt").Each(func(i int, s *goquery.Selection) {
		amt_string := strings.Split(s.Text(), "$")[1]
		amt, err := strconv.ParseFloat(amt_string, 32)
		if err != nil {
			return
		}
		amt = amount_round(amt, 2)

		switch i {
		case 0:
			data.Balances.StandardMealPlan = amt
		case 1:
			data.Balances.PlusMealPlan = amt
		case 2:
			data.Balances.Flex = amt
		}

	})

	//find recent transactions
	doc.Find(".transactions").Not(".summary").Find("li").Each(func(i int, s *goquery.Selection) {
		amt_string := strings.Split(s.Find(".transaction-amt").Text(), "$")[1]
		amt, err := strconv.ParseFloat(amt_string, 32)
		if err != nil {
			return
		}
		amt = amount_round(amt, 2)

		t := Transaction_T{}
		t.Amount = amt
		t.Date = s.Find(".transaction-date").Text()
		t.Location = s.Find(".transaction-desc").Text()
		data.Transactions.Recent = append(data.Transactions.Recent, t)
	})

	return data, true
}

func TestData() ONECardData {
	return ONECardData{Balances: Balances_T{StandardMealPlan: 100, Flex: 100, PlusMealPlan: 100}}
}

// not my functions rounds the floats for me though
func amount_round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
