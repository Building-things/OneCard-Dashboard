package uvicapi

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type Balances_T struct {
	StandardMealPlan string
	PlusMealPlan     string
	Flex             string
}
type Transaction_T struct {
	Location string
	Amount   string
	Date     string
	Credit   bool
}

type MetaData_T struct {
	StandardTotal float64
	Name          string
}
type ONECardData struct {
	Balances     Balances_T
	Transactions []Transaction_T
	Meta         MetaData_T
}

func GetOneCardData(c *http.Client) ONECardData {
	ONECARD_URL := "https://www.uvic.ca/MyCard/account/summary"
	//get onecard page
	resp, err := c.Get(ONECARD_URL)
	if err != nil {
		return ONECardData{}
	}

	//setup tokenizer
	tokenizer := html.NewTokenizer(resp.Body)
	tokens := make([]html.Token, 0)

	//save every token to an array
	//the reason for this is after i see what address in the array contains the tokesn i want
	//i can just hardcode the array indices and the code should be decently fast
	//this only breaks if UVIC updates the page
	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			return ONECardData{}
		}

		tokens = append(tokens, tokenizer.Token())
	}
	//code to refind indexes if they change
	// for i, v := range tokens {
	// 	fmt.Print(i, ": ")
	// 	fmt.Println(v)
	// }
	standard, err := strconv.ParseFloat(strings.Split(tokens[183].Data, "$")[1], 32)
	plus, err := strconv.ParseFloat(strings.Split(tokens[193].Data, "$")[1], 32)
	flex, err := strconv.ParseFloat(strings.Split(tokens[203].Data, "$")[1], 32)

	data := ONECardData{}
	data.Balances.StandardMealPlan = fmt.Sprintf("%.2f", standard)
	data.Balances.PlusMealPlan = fmt.Sprintf("%.2f", plus)
	data.Balances.Flex = fmt.Sprintf("%.2f", flex)
	data.Meta.Name = strings.Split(tokens[11].Data, "-")[0]
	data.Transactions = append(data.Transactions, get_transaction(80, tokens))
	data.Transactions = append(data.Transactions, get_transaction(96, tokens))
	data.Transactions = append(data.Transactions, get_transaction(112, tokens))
	data.Transactions = append(data.Transactions, get_transaction(128, tokens))
	data.Transactions = append(data.Transactions, get_transaction(144, tokens))
	return data
}

func get_transaction(start_index int, tokens []html.Token) Transaction_T {
	t_date := tokens[start_index].Data
	t_location := tokens[start_index+4].Data
	t_amount := tokens[start_index+8].Data
	is_credit := tokens[start_index+7].Attr
	is_debit := false
	debit_string := strings.Split(is_credit[0].Val, "-amt")
	if debit_string[1] == " debit" {
		is_debit = false
	} else if debit_string[1] == " credit" {
		is_debit = true
	}
	return Transaction_T{t_location, t_amount, t_date, is_debit}
}

func TestData() ONECardData {
	transactions := []Transaction_T{
		{Location: "Sydney", Amount: "123.0", Date: "1836 2731 2743"},
		{Location: "Berlin", Amount: "123.0", Date: "1836 2731 2743"},
		{Location: "Victoria", Amount: "123.0", Date: "1836 2731 2743"},
	}
	return ONECardData{
		Balances:     Balances_T{StandardMealPlan: "100", Flex: "100", PlusMealPlan: "100"},
		Transactions: transactions,
		Meta:         MetaData_T{StandardTotal: 1291.75},
	}
}
