package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	expectedJSON = `{"rates":{"CAD":1.3083475298,"HKD":7.786694163,"ISK":121.5816372277,"PHP":50.7253653725,"DKK":6.6981081323,"HUF":296.6287097642,"CZK":22.8718730386,"GBP":0.7632923877,"RON":4.2874562898,"SEK":9.3573926298,"IDR":13963.8751905317,"INR":71.4041065184,"BRL":4.048865776,"RUB":62.0472518605,"HRK":6.6775755402,"JPY":109.5938312562,"THB":30.1551152156,"CHF":0.9764189007,"EUR":0.8966197436,"MYR":4.1274993275,"BGN":1.7536088945,"TRY":5.957948534,"CNY":6.9958755492,"NOK":8.8386981081,"NZD":1.4933201829,"ZAR":14.045458621,"USD":1.0,"MXN":18.9176903075,"SGD":1.3523715592,"AUD":1.4342329418,"ILS":3.4619384919,"KRW":1160.6025284677,"PLN":3.8259661078},"base":"USD","date":"2019-12-27"}`
)

func TestExchangeRate(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, expectedJSON)
		}))
	defer ts.Close()
	url = ts.URL

	got, err := runExchangeRate()
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	expected := "BRL: 4.05 (USD 0.25), EUR: 0.90 (USD 1.12), GBP: 0.76 (USD 1.31)"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
