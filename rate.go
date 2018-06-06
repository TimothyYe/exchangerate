package core

// Query exchange rates for multiple currencies
func Query(from string, amount float32, to []string) map[string]string {
	helper := NewRateHelper()
	helper.FromCurrency = from
	helper.ToCurrency = to
	helper.Amount = amount

	// Start query
	helper.Query()

	// Return result
	return helper.Result
}
