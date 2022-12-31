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

func GetOneCardData(c *http.Client) ONECardData {
	return ONECardData{}
}

func TestData() ONECardData {
	return ONECardData{
		Balances:     Balances_T{StandardMealPlan: 100, Flex: 100, PlusMealPlan: 100},
		Transactions: Transactions_T{Recent: make([]Transaction_T, 3), All: make([]Transaction_T, 10)},
		Meta:         MetaData_T{StandardTotal: 1291.75},
	}
}
