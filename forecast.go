package forecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	baseURI = "https://api.forecast.io/forecast"
)

func Get(apiKey string, latitude float64, longitude float64) (*Data, error) {
	url := fmt.Sprintf("%s/%s/%f,%f", baseURI, apiKey, latitude, longitude)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	var data Data
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
