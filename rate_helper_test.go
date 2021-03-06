package exchangerate

import (
	"fmt"
	"testing"
)

func TestNewRateHelper(t *testing.T) {
	helper := NewRateHelper()

	if helper == nil {
		t.Fatal("New rate helper failed...")
	}
}

func TestQuery(t *testing.T) {
	helper := NewRateHelper()

	if helper == nil {
		t.Fatal("New rate helper failed...")
		return
	}

	helper.FromCurrency = "USD"
	helper.ToCurrency = []string{"CNY", "JPY", "AUD"}
	helper.Query()

	if len(helper.Result) != 3 {
		t.Fatal("Query result is not correct...")
	}

	for k, v := range helper.Result {
		fmt.Printf("%s %s\r\n", k, v)
	}
}

func TestQueryWithAmount(t *testing.T) {
	helper := NewRateHelper()

	if helper == nil {
		t.Fatal("New rate helper failed...")
		return
	}

	helper.FromCurrency = "USD"
	helper.Amount = 2
	helper.ToCurrency = []string{"CNY", "JPY", "AUD"}
	helper.Query()

	if len(helper.Result) != 3 {
		t.Fatal("Query result is not correct...")
	}

	for k, v := range helper.Result {
		fmt.Printf("%s %s\r\n", k, v)
	}
}
