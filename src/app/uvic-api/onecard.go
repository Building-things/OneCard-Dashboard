package uvicapi

import "net/http"

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
	return ONECardData{}
}

func TestData() ONECardData {
	transactions := []Transaction_T{
		{Location: "Sydney", Amount: 123.0, Date: "1836 2731 2743"},
		{Location: "Berlin", Amount: 123.0, Date: "1836 2731 2743"},
		{Location: "Victoria", Amount: 123.0, Date: "1836 2731 2743"},
	}
	return ONECardData{
		Balances:     Balances_T{StandardMealPlan: 100, Flex: 100, PlusMealPlan: 100},
		Transactions: transactions,
		Meta:         MetaData_T{StandardTotal: 1291.75},
	}
}
