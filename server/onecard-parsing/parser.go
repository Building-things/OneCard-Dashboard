package onecardparsing

import (
	"strconv"
	"strings"
	"sync"

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
type MetaData_T struct {
	StandardTotal float64
}
type ONECardData struct {
	Balances     Balances_T
	Transactions Transactions_T
	Meta         MetaData_T
}

func OneCardData(username string, password string) (ONECardData, bool) {
	//creates a client with authenticated cookies
	client, ok := createAuthenticatedClient(username, password)
	if !ok {
		return ONECardData{}, false
	}

	ONECARD_URL := "https://www.uvic.ca/MyCard/account/summary"

	//get onecard page
	resp, err := client.Get(ONECARD_URL)
	if err != nil {
		return ONECardData{}, false
	}

	//parse user data from the onecard page
	doc, doc_err := goquery.NewDocumentFromReader(resp.Body)
	if doc_err != nil {
		return ONECardData{}, false
	}

	var wg sync.WaitGroup
	wg.Add(3)
	balances_chan := make(chan Balances_T)
	transactions_chan := make(chan Transactions_T)
	meta_chan := make(chan MetaData_T)

	//parse balances
	go getBalances(doc, &wg, balances_chan)

	//find recent transactions
	go getTransactions(doc, &wg, transactions_chan)

	//get some meta deta about the meal plan (term dates, total ect.)
	go getMeta(&wg, meta_chan)

	data := ONECardData{Balances: <-balances_chan, Transactions: <-transactions_chan, Meta: <-meta_chan}

	wg.Wait()
	return data, true
}

func TestData() ONECardData {
	return ONECardData{
		Balances:     Balances_T{StandardMealPlan: 100, Flex: 100, PlusMealPlan: 100},
		Transactions: Transactions_T{Recent: make([]Transaction_T, 3), All: make([]Transaction_T, 10)},
		Meta:         MetaData_T{StandardTotal: 1291.75},
	}
}

func getBalances(doc *goquery.Document, wg *sync.WaitGroup, c chan Balances_T) {
	defer wg.Done()

	data := Balances_T{}
	doc.Find(".summary.transactions > ul > li > .transaction-amt").Each(func(i int, s *goquery.Selection) {
		amt_string := strings.Split(s.Text(), "$")[1]
		amt, err := strconv.ParseFloat(amt_string, 32)
		if err != nil {
			return
		}
		amt = amount_round(amt, 2)

		switch i {
		case 0:
			data.StandardMealPlan = amt
		case 1:
			data.PlusMealPlan = amt
		case 2:
			data.Flex = amt
		}

	})
	c <- data
}

func getTransactions(doc *goquery.Document, wg *sync.WaitGroup, c chan Transactions_T) {
	defer wg.Done()

	data := Transactions_T{}
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
		data.Recent = append(data.Recent, t)
	})
	data.All = data.Recent
	c <- data
}

func getMeta(wg *sync.WaitGroup, c chan MetaData_T) {
	defer wg.Done()
	c <- MetaData_T{StandardTotal: 1291.75}
}
