package core

func Query(from string, to []string, amount float32) {
	helper := NewRateHelper()
	helper.FromCurrency = from
	helper.ToCurrency = to
	helper.Amount = amount

	// Start query
	helper.Query()

	// Parse result
}
