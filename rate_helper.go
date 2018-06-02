package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/tidwall/gjson"
)

const apiAddress = "http://free.currencyconverterapi.com/api/v5/convert?q=%s_%s&compact=y"

// RateHelper struct definition
type RateHelper struct {
	FromCurrency string
	ToCurrency   []string
	Result       map[string]string
	mutex        sync.Mutex
}

// NewRateHelper creates and return a new rate helper struct
func NewRateHelper() *RateHelper {
	return &RateHelper{}
}

// SaveResult saves rate into internal map
func (helper *RateHelper) SaveResult(key, value string) {
	helper.mutex.Lock()
	helper.Result[key] = value
	helper.mutex.Unlock()
}

// Query function, queries different exchange rates and return the results
func (helper *RateHelper) Query() {
	helper.Result = map[string]string{}

	var wg sync.WaitGroup

	for _, to := range helper.ToCurrency {
		wg.Add(1)

		go func(to string) {
			resp, err := http.Get(fmt.Sprintf(apiAddress, helper.FromCurrency, to))
			if err != nil {
				helper.SaveResult(to, "N/A")
				return
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
				helper.SaveResult(to, "N/A")
				return
			}

			value := gjson.Get(string(body), fmt.Sprintf("%s_%s.val", helper.FromCurrency, to))

			fmt.Println(value.String())
			helper.SaveResult(to, value.String())
			wg.Done()
		}(to)
	}

	wg.Wait()
}
