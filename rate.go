package core

import "fmt"

// Query exchange rates for multiple currencies
func Query(from string, amount float32, to []string) {
	helper := NewRateHelper()
	helper.FromCurrency = from
	helper.ToCurrency = to
	helper.Amount = amount

	// Start query
	helper.Query()

	// Parse result
	for k, v := range helper.Result {
		fmt.Printf("%s: %s \r\n", k, v)
	}
}
