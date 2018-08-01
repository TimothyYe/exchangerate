package core

import "testing"

func TestRateQuery(t *testing.T) {
	result := Query("USD", 1, []string{"CNY", "JPY"})

	if len(result) != 2 {
		t.Fatal("Query result is not correct...")
	}

	if _, ok := result["CNY"]; !ok {
		t.Fatal("Expected result doesn't exist")
	}

	if _, ok := result["JPY"]; !ok {
		t.Fatal("Expected result doesn't exist")
	}
}
