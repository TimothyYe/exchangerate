package exchangerate

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRateQuery(t *testing.T) {
	content, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".er"))
	if err != nil {
		t.Fatal("Failed to load API key")
	}

	apiKey := strings.Replace(string(content), "\n", "", -1)
	result := Query(apiKey, "USD", 1, []string{"CNY", "JPY"})

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
