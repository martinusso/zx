package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	jsonIP8888 = `{"ipAddress": "8.8.8.8","continentCode": "NA","continentName": "North America","countryCode": "US","countryName": "United States","stateProv": "California","city": "Mountain View"}`
)

func TestIPGeolocation(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, jsonIP8888)
		}))
	defer ts.Close()
	ipURL = ts.URL

	got, err := runIP([]string{"8.8.8.8"})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	expected := "IP: 8.8.8.8\nContinent: North America (NA)\nCountry: United States (US)\nCity: Mountain View (California)"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestWhatIsMyIP(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			j := `{"ip":"187.36.173.130"}`
			fmt.Fprintln(w, j)
		}))
	defer ts.Close()
	urlWhatIsMyIP = ts.URL

	got, err := runIP([]string{"my"})
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	expected := "Your IPv4 Address Is: 187.36.173.130\n"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
