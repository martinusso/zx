package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusBadRequest {
		return nil, errors.New("400 Bad Request")
	}

	return ioutil.ReadAll(res.Body)
}

func GetJSON(url string, v interface{}) error {
	body, err := GetBody(url)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
