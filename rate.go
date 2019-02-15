package exchangerate

// Query exchange rates for multiple currencies
func Query(apiKey, from string, amount float32, to []string) map[string]string {
	helper := NewRateHelper()
	helper.APIKey = apiKey
	helper.FromCurrency = from
	helper.ToCurrency = to
	helper.Amount = amount

	// Start query
	helper.Query()

	// Return result
	return helper.Result
}
