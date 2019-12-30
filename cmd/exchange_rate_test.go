package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	expectedJSON = `{"rates":{"CAD":0.3231392697,"HKD":1.9231791306,"ISK":30.0285669996,"PHP":12.5282901876,"DKK":1.6543171601,"HUF":73.2621741923,"CZK":5.6489580796,"GBP":0.1885200523,"RON":1.058927741,"SEK":2.3111145559,"IDR":3448.836282304,"INR":17.6355825232,"BRL":1.0,"RUB":15.3246008371,"HRK":1.6492459641,"JPY":27.0677857254,"THB":7.4477932546,"CHF":0.2411586244,"EUR":0.2214496091,"MYR":1.0194211307,"BGN":0.4331111456,"TRY":1.4715105078,"CNY":1.7278605753,"NOK":2.183005957,"NZD":0.368824324,"ZAR":3.4689859822,"USD":0.2469827491,"MXN":4.6723431583,"SGD":0.3340124455,"AUD":0.3542307948,"ILS":0.8550390859,"KRW":286.6488030649,"PLN":0.9449476272},"base":"BRL","date":"2019-12-27"}`
)

func TestExchangeRate(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, expectedJSON)
		}))
	defer ts.Close()
	url = ts.URL

	got := exchangeRate()
	expected := "Dólar: 4.05, Euro: 4.52, Libra: 5.30"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
