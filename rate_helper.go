package core

import (
	"sync"
)

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
			helper.SaveResult(to, to)
			wg.Done()
		}(to)
	}

	wg.Wait()
}
