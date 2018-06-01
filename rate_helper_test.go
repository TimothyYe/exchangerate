package core

import (
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

	helper.ToCurrency = []string{"USD", "CNY", "AUD"}
	helper.Query()

	if len(helper.Result) != 3 {
		t.Fatal("Query result is not correct...")
	}
}
