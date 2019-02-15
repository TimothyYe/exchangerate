package exchangerate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

const apiAddress = "https://free.currencyconverterapi.com/api/v6/convert?q=%s_%s&compact=y&apiKey=%s"

// RateHelper struct definition
type RateHelper struct {
	APIKey       string
	Amount       float32
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
			defer wg.Done()

			resp, err := http.Get(fmt.Sprintf(apiAddress, helper.FromCurrency, to, helper.APIKey))
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

			valueStr := gjson.Get(string(body), fmt.Sprintf("%s_%s.val", helper.FromCurrency, to))

			if helper.Amount != 1 {
				value := decimal.NewFromFloat(valueStr.Float())
				result := value.Mul(decimal.NewFromFloat(float64(helper.Amount)))
				helper.SaveResult(to, result.StringFixedBank(2))
			} else {
				helper.SaveResult(to, valueStr.String())
			}
		}(to)
	}

	wg.Wait()
}
